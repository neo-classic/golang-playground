package task

import (
	"context"
	"fmt"
	"time"

	"github.com/neo-classic/golang-playground/rest/03_gin/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbContextTimeout = 5 * time.Second
)

type MongoTaskRepository struct {
	collName string
	dbName   string
	client   *mongo.Client
}

func wrapError(err error, message string) error {
	return errors.Wrapf(err, "[task_repo] %s", message)
}

func NewTaskRepository(hostPort, dbName, user, pwd, collection string) (*MongoTaskRepository, error) {
	connString := fmt.Sprintf("mongodb://%s:%s@%s/%s", user, pwd, hostPort, dbName)
	if user == "" {
		connString = fmt.Sprintf("mongodb://%s/%s", hostPort, dbName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbContextTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	ctx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoTaskRepository{
		collName: collection,
		dbName:   dbName,
		client:   client,
	}, nil
}

func (m *MongoTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	_, err := m.client.Database(m.dbName).Collection(m.collName).InsertOne(ctx, &task)
	if err != nil {
		return wrapError(err, "fail create a task")
	}

	return nil
}

func (m *MongoTaskRepository) Fetch(ctx context.Context, filter Filter) ([]*domain.Task, error) {
	var res []*domain.Task

	filters := decodeFilter(filter)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})

	cur, err := m.client.Database(m.dbName).Collection(m.collName).Find(ctx, filters, findOptions)
	if err != nil {
		return nil, wrapError(err, "fail to fetch tasks")
	}
	defer func() {
		_ = cur.Close(ctx)
	}()

	for cur.Next(ctx) {
		var t domain.Task
		err := cur.Decode(&t)
		if err != nil {
			return nil, wrapError(err, "fail to fetch task from db")
		}
		res = append(res, &t)
	}

	if err := cur.Err(); err != nil {
		return nil, wrapError(err, "fail to close a db cursor")
	}

	return res, nil
}

func (m *MongoTaskRepository) GetByGUID(ctx context.Context, guid domain.TaskGUID) (task *domain.Task, err error) {
	err = m.client.Database(m.dbName).Collection(m.collName).FindOne(ctx, bson.M{"guid": guid}).Decode(&task)
	if err != nil {
		return nil, wrapError(err, fmt.Sprintf("fail to find task with guid: %s", guid))
	}

	return task, nil
}

func (m *MongoTaskRepository) Delete(ctx context.Context, guid domain.TaskGUID) error {
	deleteResult, err := m.client.Database(m.dbName).Collection(m.collName).DeleteOne(ctx, bson.M{"guid": guid})
	if err != nil {
		return wrapError(err, "fail delete task")
	}
	if deleteResult.DeletedCount != 1 {
		return wrapError(errors.New("task not found"), "")
	}

	return nil
}

func (m *MongoTaskRepository) DeleteAll(ctx context.Context) error {
	_, err := m.client.Database(m.dbName).Collection(m.collName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return wrapError(err, "fail delete all tasks")
	}
	return nil
}

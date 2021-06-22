package task

import "go.mongodb.org/mongo-driver/bson"

type Filter struct {
	Tag     string
	DueDate string
}

func decodeFilter(f Filter) bson.M {
	var filters []bson.M

	//if f.Tag != "" {
	//	filters = append(filters, bson.M{"tags": bson.M{"$in": f.Tag}})
	//}

	if f.Tag != "" {
		filters = append(filters, bson.M{"tags": f.Tag})
	}

	if f.DueDate != "" {
		filters = append(filters, bson.M{"due": bson.M{"$gt": f.DueDate}})
	}

	if len(filters) == 1 {
		return filters[0]
	} else if len(filters) > 1 {
		return bson.M{"$and": filters}
	}

	return bson.M{}
}

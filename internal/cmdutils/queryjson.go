package cmdutils

import "github.com/itchyny/gojq"

type QueryResult struct {
	Result []interface{}
}

// QueryJson can query an unmarshalled json object with gojq.
func QueryJSON(o []interface{}, queryResult *QueryResult, queryStr string) error {
	query, err := gojq.Parse(queryStr)
	if err != nil {
		return err
	}

	iter := query.Run(o)

	for {
		value, ok := iter.Next()
		if !ok {
			break
		}

		if err, ok := value.(error); ok {
			return err
		}

		queryResult.Result = append(queryResult.Result, value)
	}

	return nil
}

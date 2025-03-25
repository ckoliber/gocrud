package service

import (
	"context"

	"github.com/ckoliber/gocrud/internal/schema"
)

type PutBulkInput[Model any] struct {
	Fields schema.Fields[Model] `query:"fields,deepObject" doc:"Entity fields" example:"[]"`
	Body   []Model
}
type PutBulkOutput[Model any] struct {
	Body []Model
}

func (s *CRUDService[Model]) PutBulk(ctx context.Context, i *PutBulkInput[Model]) (*PutBulkOutput[Model], error) {
	o := &PutBulkOutput[Model]{}

	// TODO
	// tx, err := s.repo.Transaction()
	// if err != nil {
	// 	return nil, err
	// }

	for _, model := range i.Body {
		// TODO: must change model ID
		where := schema.Where[Model]{"id": "#model.ID"}

		if err := s.hooks.PreUpdate(&i.Fields, &where, nil, nil, nil, &model); err != nil {
			// tx.Rollback()
			return nil, err
		}

		result, err := s.repo.Update(&i.Fields, &where, nil, nil, nil, &model)
		if err != nil {
			// tx.Rollback()
			return nil, err
		}

		if err := s.hooks.PostUpdate(&result); err != nil {
			// tx.Rollback()
			return nil, err
		}

		o.Body = append(o.Body, result...)
	}

	// tx.Commit()

	return o, nil
}

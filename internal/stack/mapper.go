package stack

func toStackDomain(entity stackEntity) Stack {
	return Stack{
		Id:              entity.Id,
		Name:            entity.Name,
		Description:     entity.Description,
		ApplicationName: entity.ApplicationName,
		Segment: Segment{
			Name:         entity.SegmentName,
			IsProductive: entity.SegIsProductive,
		},
		Purpose: Purpose{
			Name:         entity.PurposeName,
			Description:  entity.PurposeDescription,
			IsProductive: entity.PurpIsProductive,
		},
		CreatedBy: entity.CreatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedBy: entity.UpdatedBy,
		UpdatedAt: entity.UpdatedAt,
	}
}

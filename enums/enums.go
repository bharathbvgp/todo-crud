package enums

type PriorityType int

const (
	LowPriority PriorityType = iota + 1
    MediumPriority
    HighPriority
)

type StatusType string

const (
    NewStatus        StatusType = "New"
    InProgressStatus StatusType = "InProgress"
    CompletedStatus  StatusType = "Completed"
    CanceledStatus   StatusType = "Canceled"
)

type CategoryType string

const (
    WorkCategory    CategoryType = "Work"
    PersonalCategory CategoryType = "Personal"
    StudyCategory    CategoryType = "Study"
)
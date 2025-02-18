package jmaeew

type TypeMessage struct {
	Type string
}

type Issue struct {
	Source string
	Status string
}

type Accuracy struct {
	Epicenter string
	Depth     string
	Magnitude string
}

type MaxIntChange struct {
	String string
	Reason string
}

type WarnArea struct {
	Chiiki  string
	Shindo1 string
	Shindo2 string
	Time    string
	Type    string
	Arrive  bool
}

type JMAEEW struct {
	Type          string
	Title         string
	CodeType      string
	Issue         Issue
	EventID       string
	Serial        int
	AnnouncedTime string
	OriginTime    string
	Hypocenter    string
	Latitude      float64
	Longitude     float64
	Magunitude    float64
	Depth         int
	MaxIntensity  string
	Accuracy      Accuracy
	MaxIntChange  MaxIntChange
	WarnArea      []WarnArea
	IsSea         bool
	IsTraining    bool
	IsAssumption  bool
	IsWarn        bool
	IsFinal       bool
	IsCancel      bool
	OriginalText  string
	Pond          string
}

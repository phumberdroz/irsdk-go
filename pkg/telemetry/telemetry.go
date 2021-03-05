package telemetry

type TelemetryData struct {
	// NEW
	CarIdxBestLapNum             int
	CarIdxBestLapTime            float32
	CarIdxLapCompleted           int
	CarIdxLastLapTime            float32
	CarIdxP2P_Count              int
	CarIdxP2P_Status             bool
	CarIdxPaceFlags              int
	CarIdxPaceLine               int
	CarIdxPaceRow                int
	CarIdxQualTireCompound       int
	CarIdxQualTireCompoundLocked bool
	CarIdxTireCompound           int
	CarIdxTrackSurfaceMaterial   int
	DcStarter                    bool
	DpFastRepair                 float32
	DpFuelAddKg                  float32
	DpFuelFill                   float32
	DpLFTireChange               float32
	DpLFTireColdPress            float32
	DpLRTireChange               float32
	DpLRTireColdPress            float32
	DpRFTireChange               float32
	DpRFTireColdPress            float32
	DpRRTireChange               float32
	DpRRTireColdPress            float32
	DpWindshieldTearoff          float32
	FastRepairAvailable          int
	FastRepairUsed               int
	FrontTireSetsAvailable       int
	FrontTireSetsUsed            int
	FuelUsePerHour               float32
	HandbrakeRaw                 float32
	LFTiresAvailable             int
	LFTiresUsed                  int
	LFshockDefl_ST               float32
	LFshockVel_ST                float32
	LRTiresAvailable             int
	LRTiresUsed                  int
	LRshockDefl_ST               float32
	LRshockVel_ST                float32
	LapCompleted                 int
	LatAccel_ST                  float32
	LeftTireSetsAvailable        int
	LeftTireSetsUsed             int
	LoadNumTextures              bool
	LongAccel_ST                 float32
	ManualBoost                  bool
	ManualNoBoost                bool
	OkToReloadTextures           bool
	PaceMode                     int
	PitSvTireCompound            int
	PitchRate_ST                 float32
	PitsOpen                     bool
	PitstopActive                bool
	PlayerCarDriverIncidentCount int
	PlayerCarDryTireSetLimit     int
	PlayerCarIdx                 int
	PlayerCarInPitStall          bool
	PlayerCarMyIncidentCount     int
	PlayerCarPitSvStatus         int
	PlayerCarPowerAdjust         float32
	PlayerCarTeamIncidentCount   int
	PlayerCarTowTime             float32
	PlayerCarWeightPenalty       float32
	PlayerTireCompound           int
	PlayerTrackSurface           int
	PlayerTrackSurfaceMaterial   int
	PushToPass                   bool
	RFTiresAvailable             int
	RFTiresUsed                  int
	RFshockDefl_ST               float32
	RFshockVel_ST                float32
	RRTiresAvailable             int
	RRTiresUsed                  int
	RRshockDefl_ST               float32
	RRshockVel_ST                float32
	RearTireSetsAvailable        int
	RearTireSetsUsed             int
	RightTireSetsAvailable       int
	RightTireSetsUsed            int
	RollRate_ST                  float32
	SessionLapsRemainEx          int
	SessionTimeOfDay             float32
	SteeringWheelTorque_ST       float32
	TireLF_RumblePitch           float32
	TireLR_RumblePitch           float32
	TireRF_RumblePitch           float32
	TireRR_RumblePitch           float32
	TireSetsAvailable            int
	TireSetsUsed                 int
	TrackTempCrew                float32
	VelocityX_ST                 float32
	VelocityY_ST                 float32
	VelocityZ_ST                 float32
	VertAccel_ST                 float32
	YawNorth                     float32
	YawRate_ST                   float32

	// bools
	DriverMarker                   bool
	IsOnTrack                      bool
	IsReplayPlaying                bool
	IsDiskLoggingEnabled           bool
	IsDiskLoggingActive            bool
	CarIdxOnPitRoad                bool
	OnPitRoad                      bool
	LapDeltaToBestLap_OK           bool
	LapDeltaToOptimalLap_OK        bool
	LapDeltaToSessionBestLap_OK    bool
	LapDeltaToSessionOptimalLap_OK bool
	LapDeltaToSessionLastlLap_OK   bool
	IsOnTrackCar                   bool
	IsInGarage                     bool
	ReplayPlaySlowMotion           bool

	// ints
	SessionNum                int
	SessionState              int
	SessionUniqueID           int
	SessionLapsRemain         int
	RadioTransmitCarIdx       int
	RadioTransmitRadioIdx     int
	RadioTransmitFrequencyIdx int
	ReplayFrameNum            int
	ReplayFrameNumEnd         int
	CarIdxLap                 int
	CarIdxTrackSurface        int
	CarIdxGear                int
	Gear                      int
	Lap                       int
	RaceLaps                  int
	LapBestLap                int
	CamCarIdx                 int
	CamCameraNumber           int
	CamGroupNumber            int
	ReplayPlaySpeed           int
	ReplaySessionNum          int
	DisplayUnits              int
	PlayerCarPosition         int
	PlayerCarClassPosition    int
	CarIdxPosition            int
	CarIdxClassPosition       int
	LapLasNLapSeq             int
	LapBestNLapLap            int
	EnterExitReset            int
	DCLapStatus               int
	DCDriversSoFar            int
	SessionTick               int

	// Only used in disk based telemetry
	WeatherType int
	Skies       int

	// bitfields
	SessionFlags   string // Todo make map[string]bool
	CamCameraState string // Todo map[string]bool
	EngineWarnings string // Todo map[string]bool
	CarLeftRight   string // Todo map[string]bool

	// Only used in disk based telemetry data
	PitSvFlags string // Todo make map[string]bool

	// floats
	FrameRate                       float32
	CpuUsageBG                      float32
	CarIdxLapDistPct                float32
	CarIdxSteer                     float32
	CarIdxRPM                       float32
	SteeringWheelAngle              float32
	Throttle                        float32
	Brake                           float32
	Clutch                          float32
	RPM                             float32
	LapDist                         float32
	LapDistPct                      float32
	LapBestLapTime                  float32
	LapLastLapTime                  float32
	LapCurrentLapTime               float32
	LapDeltaToBestLap               float32
	LapDeltaToBestLap_DD            float32
	LapDeltaToOptimalLap            float32
	LapDeltaToOptimalLap_DD         float32
	LapDeltaToSessionBestLap        float32
	LapDeltaToSessionBestLap_DD     float32
	LapDeltaToSessionOptimalLap     float32
	LapDeltaToSessionOptimalLap_DD  float32
	LapDeltaToSessionLastlLap       float32
	LapDeltaToSessionLastlLap_DD    float32
	LongAccel                       float32
	LatAccel                        float32
	VertAccel                       float32
	RollRate                        float32
	PitchRate                       float32
	YawRate                         float32
	Speed                           float32
	VelocityX                       float32
	VelocityY                       float32
	VelocityZ                       float32
	Yaw                             float32
	Pitch                           float32
	Roll                            float32
	PitRepairLeft                   float32
	PitOptRepairLeft                float32
	SteeringWheelTorque             float32
	SteeringWheelPctTorque          float32
	SteeringWheelPctTorqueSign      float32
	SteeringWheelPctTorqueSignStops float32
	SteeringWheelPctDamper          float32
	SteeringWheelAngleMax           float32
	ShiftIndicatorPct               float32
	ShiftPowerPct                   float32
	ShiftGrindRPM                   float32
	ThrottleRaw                     float32
	BrakeRaw                        float32
	SteeringWheelPeakForceNm        float32
	FuelLevel                       float32
	FuelLevelPct                    float32
	WaterTemp                       float32
	WaterLevel                      float32
	FuelPress                       float32
	OilTemp                         float32
	OilPress                        float32
	OilLevel                        float32
	Voltage                         float32
	ManifoldPress                   float32
	RRbrakeLinePress                float32
	RRcoldPressure                  float32
	RRtempCL                        float32
	RRtempCM                        float32
	RRtempCR                        float32
	RRwearL                         float32
	RRwearM                         float32
	RRwearR                         float32
	LRbrakeLinePress                float32
	LRcoldPressure                  float32
	LRtempCL                        float32
	LRtempCM                        float32
	LRtempCR                        float32
	LRwearL                         float32
	LRwearM                         float32
	LRwearR                         float32
	RFbrakeLinePress                float32
	RFcoldPressure                  float32
	RFtempCL                        float32
	RFtempCM                        float32
	RFtempCR                        float32
	RFwearL                         float32
	RFwearM                         float32
	RFwearR                         float32
	LFbrakeLinePress                float32
	LFcoldPressure                  float32
	LFtempCL                        float32
	LFtempCM                        float32
	LFtempCR                        float32
	LFwearL                         float32
	LFwearM                         float32
	LFwearR                         float32
	RRshockDefl                     float32
	RRshockVel                      float32
	LRshockDefl                     float32
	LRshockVel                      float32
	RFshockDefl                     float32
	RFshockVel                      float32
	LFshockDefl                     float32
	LFshockVel                      float32
	CarIdxF2Time                    float32
	CarIdxEstTime                   float32
	LapLastNLapTime                 float32
	BrakeLinePresse                 float32
	DcBrakeBias                     float32
	LapBestNLapTime                 float32

	// Only used in disk based telemetry
	Alt               float32
	TrackTemp         float32
	AirTemp           float32
	AirDensity        float32
	AirPressure       float32
	WindVel           float32
	WindDir           float32
	RelativeHumidity  float32
	LRtempL           float32
	LRtempM           float32
	LRtempR           float32
	RFspeed           float32
	RFpressure        float32
	DcABS             float32
	DcThrottleShape   float32
	DcFuelMixture     float32
	RRspeed           float32
	RRpressure        float32
	RRtempL           float32
	RRtempM           float32
	RRtempR           float32
	LRspeed           float32
	LRpressure        float32
	RFtempL           float32
	RFtempM           float32
	RFtempR           float32
	LFspeed           float32
	LFpressure        float32
	LFtempL           float32
	LFtempM           float32
	LFtempR           float32
	LFrideHeight      float32
	RFrideHeight      float32
	LRrideHeight      float32
	RRrideHeight      float32
	CFSRrideHeight    float32
	PitSvLRP          float32
	PitSvRRP          float32
	FogLevel          float32
	DcTractionControl float32
	PitSvLFP          float32
	PitSvRFP          float32
	PitSvFuel         float32

	// Doubles
	SessionTime       float64
	SessionTimeRemain float64
	ReplaySessionTime float64

	// Only used in disk based telemetry
	Lat float64
	Lon float64
}

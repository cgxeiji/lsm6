package lsm6

// Register addresses
const (
	FuncCfgAccess = 0x01

	FIFOCtrl1  = 0x06
	FIFOCtrl2  = 0x07
	FIFOCtrl3  = 0x08
	FIFOCtrl4  = 0x09
	FIFOCtrl5  = 0x0A
	OrientCfgG = 0x0B

	Int1Ctrl = 0x0D
	Int2Ctrl = 0x0E
	WhoAmI   = 0x0F
	Ctrl1XL  = 0x10
	Ctrl2G   = 0x11
	Ctrl3C   = 0x12
	Ctrl4C   = 0x13
	Ctrl5C   = 0x14
	Ctrl6C   = 0x15
	Ctrl7G   = 0x16
	Ctrl8XL  = 0x17
	Ctrl9XL  = 0x18
	Ctrl10C  = 0x19

	WakeUpSrc = 0x1B
	TapSrc    = 0x1C
	D6DSrc    = 0x1D
	StatusReg = 0x1E

	OutTempL = 0x20
	OutTempH = 0x21
	OutXLG   = 0x22
	OutXHG   = 0x23
	OutYLG   = 0x24
	OutYHG   = 0x25
	OutZLG   = 0x26
	OutZHG   = 0x27
	OutXLXL  = 0x28
	OutXHXL  = 0x29
	OutYLXL  = 0x2A
	OutYHXL  = 0x2B
	OutZLXL  = 0x2C
	OutZHXL  = 0x2D

	FIFOStatus1   = 0x3A
	FIFOStatus2   = 0x3B
	FIFOStatus3   = 0x3C
	FIFOStatus4   = 0x3D
	FIFODataOutL  = 0x3E
	FIFODataOutH  = 0x3F
	Timestamp0Reg = 0x40
	Timestamp1Reg = 0x41
	Timestamp2Reg = 0x42

	StepTimestampL = 0x49
	StepTimestampH = 0x4A
	StepCounterL   = 0x4B
	StepCounterH   = 0x4C

	FuncSrc = 0x53

	TapCfg    = 0x58
	TapThs6D  = 0x59
	IntDur2   = 0x5A
	WakeUpThs = 0x5B
	WakeUpDur = 0x5C
	FreeFall  = 0x5D
	MD1Cfg    = 0x5E
	MD2Cfg    = 0x5F
)

// Device constants
const (
	Addr = 0x6B
)

const (
	maxADC  = (1 << 16) - 1
	halfADC = (1 << 15) - 1
)

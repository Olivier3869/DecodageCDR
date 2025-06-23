package symp_home

import "local/structures/defStructs"

// T2sympsssas
var T2sympsssas = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympsssasdc},
}

// T3sympsssasdc
var T3sympsssasdc = map[int]defStructs.Champ{
	0x30: {"MediaCaller", -1, -1, &T4sympsssasdcMediaCaller},
	0x31: {"MediaDestUser", -1, -1, &T4sympsssasdcMediaDestUser},
	0x21: {"SDPMedia1", -1, -1, &T4sympsssasdcSDPMedia},
	0x22: {"SDPMedia2", -1, -1, &T4sympsssasdcSDPMedia},
	0x23: {"SDPMedia3", -1, -1, &T4sympsssasdcSDPMedia},
	0x24: {"SDPMedia4", -1, -1, &T4sympsssasdcSDPMedia},
	0x25: {"SDPMedia5", -1, -1, &T4sympsssasdcSDPMedia},
	0x40: {"CallingNb", -1, 1, nil},
	0x04: {"OrigCalledNb", -1, 1, nil},
	0x05: {"StopSecretId", -1, 1, nil},
	0x06: {"CallPartyAdd", -1, 1, nil},
	0x07: {"ImsChargingId", -1, 1, nil},
	0x0c: {"StopSecretAction", -1, -1, nil},
	0x01: {"Date", -1, -1, nil},
	0x08: {"SessionEndTime", -1, -1, nil},
	0x13: {"Duree", -1, -1, nil},
	0x0b: {"DatePartyCalled", -1, -1, nil},
	0x0d: {"DureePartyCalled", -1, -1, nil},
	0x0e: {"DateCall", -1, -1, nil},
	0x50: {"ReleasingParty", -1, -1, nil},
	0x51: {"TerminationCause", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T4sympsssasdcMediaCaller
var T4sympsssasdcMediaCaller = map[int]defStructs.Champ{
	0x01: {"CallerMediaName", -1, 1, nil},
	0x02: {"CallerListOfMediaDesc1", -1, 1, nil},
	0x03: {"CallerListOfMediaDesc2", -1, 1, nil},
	0x04: {"CallerListOfMediaDesc3", -1, 1, nil},
	0x05: {"CallerListOfMediaDesc4", -1, 1, nil},
	0x06: {"CallerListOfMediaDesc5", -1, 1, nil},
}

// T4sympsssasdcMediaDestUser
var T4sympsssasdcMediaDestUser = map[int]defStructs.Champ{
	0x01: {"DestMediaName", -1, 1, nil},
	0x02: {"CallerListOfMediaDesc1", -1, 1, nil},
	0x03: {"CallerListOfMediaDesc2", -1, 1, nil},
	0x04: {"CallerListOfMediaDesc3", -1, 1, nil},
	0x05: {"CallerListOfMediaDesc4", -1, 1, nil},
	0x06: {"CallerListOfMediaDesc5", -1, 1, nil},
}

// T4sympsssasdcSDPMedia
var T4sympsssasdcSDPMedia = map[int]defStructs.Champ{
	0x01: {"MediaDuration", -1, -1, nil},
	0x02: {"MediaTime", -1, -1, nil},
	0x03: {"MediaName1", -1, 1, nil},
	0x04: {"MediaName2", -1, 1, nil},
	0x05: {"MediaName3", -1, 1, nil},
	0x06: {"MediaName4", -1, 1, nil},
	0x07: {"MediaName5", -1, 1, nil},
}

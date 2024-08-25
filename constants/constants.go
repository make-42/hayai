package constants

// https://earthquake.usgs.gov/earthquakes/eventpage/us6000d3zh/dyfi/intensity-vs-distance ignored since too tame
// https://earthquake.usgs.gov/earthquakes/eventpage/us6000nith/dyfi/intensity-vs-distance used
const MagnitudeDecreasePerKM = 0.00814867762 // 0.01159528707
var TestMessage = []byte(`{"type":"jma_eew","Title": "緊急地震速報（予報）", "CodeType": "Ｍ、最大予測震度及び主要動到達予測時刻の緊急地震速報", "Issue": {"Source": "東京", "Status": "通常"}, "EventID": "20240824093249", "Serial": 5, "AnnouncedTime": "2024/08/24 09:33:35", "OriginTime": "2024/08/24 09:32:40", "Hypocenter": "四国沖", "Latitude": 32.9, "Longitude": 134.3, "Magunitude": 4.0, "Depth": 10, "MaxIntensity": "2", "Accuracy": {"Epicenter": "IPF 法（5 点以上）", "Depth": "IPF 法（5 点以上）", "Magnitude": "防災科研システム"}, "MaxIntChange": {"String": "ほとんど変化なし", "Reason": "不明、未設定時、キャンセル時"}, "WarnArea": [ {"Chiiki": "Test area", "Shindo1":"1","Shindo2":"2","Time":"2024/08/24 09:33:35","Type":"警報","Arrive":false} ], "isSea": true, "isTraining": false, "isAssumption": false, "isWarn": false, "isFinal": true, "isCancel": false, "OriginalText": "37 03 00 240824093335 C11 240824093240 ND20240824093249 NCN905 JD////////////// JN/// 902 N329 E1343 010 40 02 RK44209 RT10/// RC0//// 9999=", "Pond": "41"}`)

const WSHost = "ws-api.wolfx.jp"

const OpenURLA = "https://bs.wolfx.jp/newJMAEQList/"
const OpenURLB = "https://smi.lmoniexp.bosai.go.jp/"

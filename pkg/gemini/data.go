package gemini

/*
Example Canidate Json:
{"Candidates":[{"Index":0,"Content":{"Parts":["```json\n[\n  {\n    \"ticket_id\": 1001,\n    \"date_created\": \"2024-04-15T10:35:00Z\",\n    \"status\": \"closed\",\n    \"priority\": \"high\",\n    \"channel\": \"email\",\n    \"customer_name\": \"Alice Johnson\",\n    \"company\": \"Acme Corp\",\n    \"environment\": \"production\",\n    \"k8s_version\": \"1.23.5\",\n    \"cloud_provider\": \"AWS\",\n    \"summary\": \"Pod CrashLoopBackOff Error\",\n    \"description\": \"Multiple pods in our production cluster are stuck in CrashLoopBackOff.  We're seeing errors related to imagePullBackOff and OOMKilled.\",\n    \"resolution\": null,\n    \"tags\": [\"pod\", \"crashloopbackoff\", \"imagepullbackoff\", \"oomkilled\", \"deployment\"]\n  }\n]\n```"],"Role":"model"},"FinishReason":1,"SafetyRatings":[{"Category":9,"Probability":1,"Blocked":false},{"Category":8,"Probability":1,"Blocked":false},{"Category":7,"Probability":1,"Blocked":false},{"Category":10,"Probability":1,"Blocked":false}],"CitationMetadata":null,"TokenCount":0}],"PromptFeedback":null,"UsageMetadata":{"PromptTokenCount":244,"CandidatesTokenCount":220,"TotalTokenCount":464}}
*/

type Candidates struct {
	Candidates []Candidate `json:"Candidates"`
}

type Candidate struct {
	Content          Content        `json:"Content"`
	FinishReason     int            `json:"FinishReason"`
	SafetyRatings    []SafetyRating `json:"SafetyRatings"`
	CitationMetadata interface{}    `json:"CitationMetadata"`
	TokenCount       int            `json:"TokenCount"`
}

type Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}

type SafetyRating struct {
	Category    int  `json:"Category"`
	Probability int  `json:"Probability"`
	Blocked     bool `json:"Blocked"`
}

/*
Exmaple of json response from generate summary
[{"ticket_id":1001,"date_created":"2024-04-15T10:35:00Z","status":"closed","priority":"high","channel":"email","customer_name":"AliceJohnson","company":"AcmeCorp","environment":"production","k8s_version":"1.23.5","cloud_provider":"AWS","summary":"PodCrashLoopBackOffError","description":"MultiplepodsinourproductionclusterarestuckinCrashLoopBackOff.We'reseeingerrorsrelatedtoimagePullBackOffandOOMKilled.","resolution":"Updateddeploymenttousecorrectimagetag,adjustedresourcelimitsforpods,monitoredforstability.","tags":["pod","crashloopbackoff","imagepullbackoff","oomkilled","deployment"]}]
*/
// JSON response struct from generate summary
type Ticket struct {
	TicketID      int      `json:"ticket_id"`
	DateCreated   string   `json:"date_created"`
	Status        string   `json:"status"`
	Priority      string   `json:"priority"`
	Channel       string   `json:"channel"`
	CustomerName  string   `json:"customer_name"`
	Company       string   `json:"company"`
	Environment   string   `json:"environment"`
	K8sVersion    string   `json:"k8s_version"`
	CloudProvider string   `json:"cloud_provider"`
	Summary       string   `json:"summary"`
	Description   string   `json:"description"`
	Resolution    string   `json:"resolution"`
	Tags          []string `json:"tags"`
}

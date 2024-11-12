package main

import "time"

const (
	Trump  string = "8639"
	Harris string = "64984"
)

// Method to get the name of the candidate based on the ID
func (c *ElectionResults) Who0() string {

	switch c.US0.Candidates[0].CandidateID {
	case Trump:
		return "Trump"
	case Harris:
		return "Harris"
	default:
		return "Unknown Candidate"
	}
}

func (c *ElectionResults) Who1() string {

	switch c.US0.Candidates[1].CandidateID {
	case Trump:
		return "Trump"
	case Harris:
		return "Harris"
	default:
		return "Unknown Candidate"
	}
}

type ElectionResults struct {
	AK0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105AK0"`
	AL0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105AL0"`
	AR0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105AR0"`
	AZ0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105AZ0"`
	CA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105CA0"`
	CO0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105CO0"`
	CT0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105CT0"`
	DC0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105DC0"`
	DE0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105DE0"`
	FL0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			Avotes      int     `json:"avotes"`
			VotePct     float64 `json:"votePct"`
			ColorIndex  int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105FL0"`
	GA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID  string  `json:"candidateID"`
			VoteCount    int     `json:"voteCount"`
			Avotes       int     `json:"avotes"`
			AdvanceTotal int     `json:"advanceTotal"`
			VotePct      float64 `json:"votePct"`
			ColorIndex   int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105GA0"`
	HI0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105HI0"`
	IA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105IA0"`
	ID0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105ID0"`
	IL0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105IL0"`
	IN0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID  string    `json:"candidateID"`
			VoteCount    int       `json:"voteCount"`
			Avotes       int       `json:"avotes"`
			Winner       string    `json:"winner,omitempty"`
			ElectWon     int       `json:"electWon,omitempty"`
			AdvanceTotal int       `json:"advanceTotal"`
			VotePct      float64   `json:"votePct"`
			CalledAt     time.Time `json:"calledAt,omitempty"`
			ColorIndex   int       `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105IN0"`
	KS0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105KS0"`
	KY0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID  string    `json:"candidateID"`
			VoteCount    int       `json:"voteCount"`
			Avotes       int       `json:"avotes"`
			Winner       string    `json:"winner,omitempty"`
			ElectWon     int       `json:"electWon,omitempty"`
			AdvanceTotal int       `json:"advanceTotal"`
			VotePct      float64   `json:"votePct"`
			CalledAt     time.Time `json:"calledAt,omitempty"`
			ColorIndex   int       `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105KY0"`
	LA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105LA0"`
	MA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MA0"`
	MD0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MD0"`
	ME0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105ME0"`
	ME82970 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105ME82970"`
	ME82971 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105ME82971"`
	MI0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MI0"`
	MN0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MN0"`
	MO0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MO0"`
	MS0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MS0"`
	MT0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105MT0"`
	NC0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID  string  `json:"candidateID"`
			VoteCount    int     `json:"voteCount"`
			Avotes       int     `json:"avotes"`
			AdvanceTotal int     `json:"advanceTotal"`
			VotePct      float64 `json:"votePct"`
			ColorIndex   int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NC0"`
	ND0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105ND0"`
	NE0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NE0"`
	NE82966 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NE82966"`
	NE82967 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NE82967"`
	NE82968 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NE82968"`
	NH0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
			ColorIndex  int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NH0"`
	NJ0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NJ0"`
	NM0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NM0"`
	NV0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NV0"`
	NY0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105NY0"`
	OH0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			Avotes      int     `json:"avotes"`
			VotePct     float64 `json:"votePct"`
			ColorIndex  int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105OH0"`
	OK0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105OK0"`
	OR0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105OR0"`
	PA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105PA0"`
	RI0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105RI0"`
	SC0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			Avotes      int     `json:"avotes"`
			VotePct     float64 `json:"votePct"`
			ColorIndex  int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105SC0"`
	SD0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105SD0"`
	TN0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105TN0"`
	TX0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105TX0"`
	US0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			ElectWon    int     `json:"electWon,omitempty"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
		Eevp        int    `json:"eevp"`
	} `json:"20241105US0"`
	UT0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105UT0"`
	VA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID  string  `json:"candidateID"`
			VoteCount    int     `json:"voteCount"`
			Avotes       int     `json:"avotes"`
			AdvanceTotal int     `json:"advanceTotal"`
			VotePct      float64 `json:"votePct"`
			ColorIndex   int     `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105VA0"`
	VT0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string    `json:"candidateID"`
			VoteCount   int       `json:"voteCount"`
			Winner      string    `json:"winner,omitempty"`
			ElectWon    int       `json:"electWon,omitempty"`
			VotePct     float64   `json:"votePct"`
			CalledAt    time.Time `json:"calledAt,omitempty"`
			ColorIndex  int       `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105VT0"`
	WA0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105WA0"`
	WI0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105WI0"`
	WV0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string    `json:"candidateID"`
			VoteCount   int       `json:"voteCount"`
			Winner      string    `json:"winner,omitempty"`
			ElectWon    int       `json:"electWon,omitempty"`
			VotePct     float64   `json:"votePct"`
			CalledAt    time.Time `json:"calledAt,omitempty"`
			ColorIndex  int       `json:"colorIndex,omitempty"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105WV0"`
	WY0 struct {
		StatePostal           string    `json:"statePostal"`
		StateName             string    `json:"stateName"`
		ReportingunitID       string    `json:"reportingunitID"`
		ReportingunitLevel    int       `json:"reportingunitLevel"`
		PollClosingTime       time.Time `json:"pollClosingTime"`
		Level                 string    `json:"level"`
		ElectTotal            int       `json:"electTotal"`
		LastUpdated           time.Time `json:"lastUpdated"`
		PrecinctsReporting    int       `json:"precinctsReporting"`
		Eevp                  float64   `json:"eevp"`
		PrecinctsTotal        int       `json:"precinctsTotal"`
		PrecinctsReportingPct float64   `json:"precinctsReportingPct"`
		Candidates            []struct {
			CandidateID string  `json:"candidateID"`
			VoteCount   int     `json:"voteCount"`
			VotePct     float64 `json:"votePct"`
		} `json:"candidates"`
		Parameters struct {
			Vote struct {
				Expected struct {
					Actual  int     `json:"actual"`
					EevpCap float64 `json:"eevpCap"`
				} `json:"expected"`
				Total      int `json:"total"`
				Registered int `json:"registered"`
			} `json:"vote"`
			PrecinctsReportingActual       int `json:"precinctsReportingActual"`
			PrecinctsReportingProportional int `json:"precinctsReportingProportional"`
		} `json:"parameters"`
		RaceID      string `json:"raceID"`
		Test        bool   `json:"test"`
		ResultsType string `json:"resultsType"`
	} `json:"20241105WY0"`
}

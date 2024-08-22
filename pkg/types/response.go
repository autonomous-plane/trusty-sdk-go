// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "time"

type Time time.Time

const timeLayout = "2006-01-02T15:04:05.000000"

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		*t = Time(time.Time{})
		return
	}

	var nt time.Time
	if strings.Contains(s, "Z") {
		nt, err = time.Parse(time.RFC3339, s)
	} else {
		nt, err = time.Parse(timeLayout, s)
	}
	*t = Time(nt)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Time) String() string {
	td := time.Time(*t)
	return fmt.Sprintf("%q", td.Format(timeLayout))
}

// Reply is the response from the package report API
type Reply struct {
	PackageName             string           `json:"package_name"`
	PackageType             string           `json:"package_type"`
	PackageVersion          string           `json:"package_version"`
	Status                  string           `json:"status"`
	Summary                 ScoreSummary     `json:"summary"`
	Provenance              *Provenance      `json:"provenance"`
	Activity                *Activity        `json:"activity"`
	Typosquatting           *Typosquatting   `json:"typosquatting"`
	Alternatives            AlternativesList `json:"alternatives"`
	PackageData             PackageData      `json:"package_data"`
	SameOriginPackagesCount int              `json:"same_origin_packages_count"`
	SimilarPackageNames     []Alternative    `json:"similar_package_names"`
}

const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusComplete   = "complete"
)

// Activity captures a package's activity score
type Activity struct {
	Score       float64             `json:"score"`
	Description ActivityDescription `json:"description"`
	UpdatedAt   *Time               `json:"updated_at"`
}

// ActivityDescription captures the fields of the activuty score
type ActivityDescription struct {
	Repository float64 `json:"repo"`
	User       float64 `json:"user"`
	// UpdatedAt  *time.Time `json:"updated_at"`
}

// Typosquatting score for the package's name
type Typosquatting struct {
	Score       float64                  `json:"score"`
	Description TyposquattingDescription `json:"description"`
	UpdatedAt   *Time                    `json:"updated_at"`
}

// TyposquattingDescription captures the dat details of the typosquatting score
type TyposquattingDescription struct {
	TotalSimilarNames int `json:"total_similar_names"`
}

// Alternative is an alternative package returned from the package intelligence API
type Alternative struct {
	ID              string  `json:"id"`
	IsMalicious     bool    `json:"is_malicious"`
	PackageName     string  `json:"package_name"`
	PackageType     string  `json:"package_type"`
	PackageVersion  string  `json:"package_version"`
	Score           float64 `json:"score"`
	PackageNameURL  string
	RepoDescription string      `json:"repo_description"`
	Provenance      *Provenance `json:"provenance"`
}

// AlternativesList is the alternatives block in the trusty API response
type AlternativesList struct {
	Status   string        `json:"status"`
	Packages []Alternative `json:"packages"`
}

// ScoreSummary is the summary score returned from the package intelligence API
type ScoreSummary struct {
	Score       *float64       `json:"score"`
	Description map[string]any `json:"description"`
	UpdatedAt   *Time          `json:"updated_at"`
}

const (
	SummaryKeyActivity      = "activity"      //float64
	SummaryKeyActivityRepo  = "activity_repo" //float64
	SummaryKeyActivityUser  = "activity_user" //float64
	SummaryKeyFrom          = "from"          //string
	SummaryKeyMalicious     = "malicious"     //bool
	SummaryKeyProvenance    = "provenance"    //float64
	SummaryKeyTrustSummary  = "trust-summary" //float64
	SummaryKeyTyposquatting = "typosquatting" //float64
)

type ScoreSummaryDescription struct {
	Activity      float64 `json:"activity"`
	ActivityRepo  float64 `json:"activity_repo"`
	ActivityUser  float64 `json:"activity_user"`
	From          string  `json:"from"`
	Malicious     bool    `json:"malicious"`
	Provenance    float64 `json:"provenance"`
	TrustSummary  float64 `json:"trust-summary"`
	Typosquatting float64 `json:"typosquatting"`
}

// PackageData contains the data about the queried package
type PackageData struct {
	Archived         bool   `json:"archived"`
	Author           string `json:"author"`
	AuthorEmail      string `json:"author_email"`
	ContributorCount int    `json:"contributor_count"`
	Contributors     []User `json:"contributors"`

	DefaultBranch   string         `json:"default_branch"`
	Disabled        bool           `json:"disabled"`
	Followers       int            `json:"followers"`
	Following       int            `json:"following"`
	ForksCount      int            `json:"forks_count"`
	HasDownloads    bool           `json:"has_downloads"`
	HasIssues       bool           `json:"has_issues"`
	HasProjects     bool           `json:"has_projects"`
	HomePage        *string        `json:"home_page"`
	ID              string         `json:"id"`
	Deprecated      bool           `json:"is_deprecated"`
	LastUpdate      *Time          `json:"last_update"`
	Malicious       *MaliciousData `json:"malicious"`
	Name            string         `json:"name"`
	OpenIssuesCount int            `json:"open_issues_count"`
	Origin          string         `json:"origin"`
	Owner           User           `json:"owner"`

	PackageDescription string         `json:"package_description"`
	PublicGists        int            `json:"public_gists"`
	PublicRepos        int            `json:"public_repos"`
	RepoDescription    string         `json:"repo_description"`
	RepositoryID       string         `json:"repository_id"`
	RepositoryName     string         `json:"repository_name"`
	Scores             map[string]any `json:"scores"` //need more data
	StargazersCount    int            `json:"stargazers_count"`
	Status             string         `json:"status"`
	StatusCode         *string        `json:"status_code"` // need more data
	Type               string         `json:"type"`
	Version            string         `json:"version"`
	VersionDate        *Time          `json:"version_date"`
	Visibility         string         `json:"visibility"`
	WatchersCount      int            `json:"watchers_count"`
}

const (
	OriginOK          = "ok"
	VisibilityPublic  = "public"
	VisibilityPrivate = "private"
)

type User struct {
	Author          string         `json:"author"`
	AuthorEmail     string         `json:"author_email"`
	AvatarURL       string         `json:"avatar_url"`
	Blog            *string        `json:"blog"`
	Company         *string        `json:"company"`
	Email           string         `json:"email"`
	Followers       int            `json:"followers"`
	Following       int            `json:"following"`
	GravatarID      string         `json:"gravatar_id"`
	Hireable        bool           `json:"hireable"`
	HtmlURL         *string        `json:"html_url"`
	ID              string         `json:"id"`
	Location        *string        `json:"location"`
	Login           string         `json:"login"`
	PublicGists     *int           `json:"public_gists"`
	PublicRepos     int            `json:"public_repos"`
	Scores          map[string]any `json:"scores"` //todo: need more data
	TwitterUsername *string        `json:"twitter_username"`
	URL             string         `json:"url"`
}

// MaliciousData contains the security details when a dependency is malicious
type MaliciousData struct {
	Summary   string `json:"summary"`
	Details   string `json:"details"`
	Published *Time  `json:"published"` //regular time values: "2024-01-16T23:40:53Z"
	Modified  *Time  `json:"modified"`  //regular time values: "2024-01-18T03:34:20Z",
	Source    string `json:"source"`
}

// Provenance has the package's provenance score and provenance type components
type Provenance struct {
	Score       float64               `json:"score"`
	Description ProvenanceDescription `json:"description"`
	UpdatedAt   *Time                 `json:"updated_at"`
}

// ProvenanceDescription contians the provenance types
type ProvenanceDescription struct {
	Historical HistoricalProvenance `json:"hp"`
	Sigstore   SigstoreProvenance   `json:"sigstore"`
}

// HistoricalProvenance has the historical provenance components from a package
type HistoricalProvenance struct {
	Tags     float64        `json:"tags"`
	Common   float64        `json:"common"`
	Overlap  float64        `json:"overlap"`
	Versions float64        `json:"versions"`
	OverTime map[string]any `json:"over_time"` //need more data
}

// SigstoreProvenance has the sigstore certificate data when a package was signed
// using a github actions workflow
type SigstoreProvenance struct {
	Issuer           string `json:"issuer"`
	Workflow         string `json:"workflow"`
	SourceRepository string `json:"source_repo"`
	TokenIssuer      string `json:"token_issuer"`
	Transparency     string `json:"transparency"`
}

package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
	"github.com/blevesearch/bleve/v2/analysis/token/keyword"
	"github.com/blevesearch/bleve/v2/mapping"
	log "github.com/sirupsen/logrus"
)

type Article struct {
	Author       string `json:"author"`
	Published    string `json:"publishdate"`
	ReadingTime  int    `json:"readingtime"`
	Section      string `json:"section"`
	Image        string `json:"image"`
	Premium      bool   `json:"premium"`
	RelPermalink string `json:"relpermalink"`
	WordCount    int    `json:"wordcount"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
}

type Service struct {
	index bleve.Index
}

func New() Service {
	mapping, err := buildIndexMapping()
	if err != nil {
		log.Error(err)
	}
	index, err := bleve.New("search.bleve", mapping)
	if err != nil {
		log.Error(err)
	}
	resp, err := http.Get("https://tutorialedge.net/algolia.json")
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var articles []Article

	err = json.Unmarshal(responseData, &articles)
	if err != nil {
		log.Error(err)
	}

	for _, article := range articles {
		err := index.Index(article.Title, article)
		if err != nil {
			log.Error(err)
		}
	}

	return Service{
		index: index,
	}
}

func (s Service) Search(searchString string) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(searchString)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := s.index.Search(search)
	if err != nil {
		log.Error(err)
	}

	return searchResults, nil
}

func buildIndexMapping() (mapping.IndexMapping, error) {

	// a generic reusable mapping for english text
	englishTextFieldMapping := bleve.NewTextFieldMapping()
	englishTextFieldMapping.Analyzer = en.AnalyzerName

	// a generic reusable mapping for keyword text
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword.Name

	postMapping := bleve.NewDocumentMapping()
	postMapping.AddFieldMappingsAt("title", englishTextFieldMapping)
	postMapping.AddFieldMappingsAt("summary", englishTextFieldMapping)

	breweryMapping := bleve.NewDocumentMapping()
	breweryMapping.AddFieldMappingsAt("title", englishTextFieldMapping)
	breweryMapping.AddFieldMappingsAt("summary", englishTextFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("post", postMapping)

	indexMapping.TypeField = "type"
	indexMapping.DefaultAnalyzer = "en"

	return indexMapping, nil
}

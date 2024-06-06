package parser

type serviceApi interface {
	parseCsv(string) error
}
type downloaderApi interface {
	sendData(*CSVData) error
}

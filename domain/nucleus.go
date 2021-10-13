package domain

type Nucleus struct {
	Number int
	Mass   int
}

type Table struct {
	Groups map[string]Group
}

type Group struct {
	nucl Nucleus
	data Data
}

type Data struct {
}

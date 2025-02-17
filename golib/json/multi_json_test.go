package json

type Struct1 struct {
	A uint64
}

type Struct2 struct {
	B string
}

type StructX interface{}

func (s *StructX) MarshalJSON() ([]byte, error) {

}



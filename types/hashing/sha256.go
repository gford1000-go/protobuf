package hashing

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"sort"
)

type sha256Hasher struct {
	s []byte
}

// GetHashType returns the HashType of the algoritm being used
func (s *sha256Hasher) GetHashType() HashType {
	return SHA256
}

// copyBytes is a helper function to copy a byte slice
func (s *sha256Hasher) copyBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

// GetSalt returns the salt used by this instance
func (s *sha256Hasher) GetSalt() []byte {
	return s.copyBytes(s.s)
}

type stringizeMapEntry struct {
	k string
	v string
}

func (sme *stringizeMapEntry) String() string {
	return fmt.Sprintf("(k %v, v %v)", sme.k, sme.v)
}

type stringizeMapSlice struct {
	data []*stringizeMapEntry
}

func (sms *stringizeMapSlice) Len() int {
	return len(sms.data)
}

func (sms *stringizeMapSlice) Less(i, j int) bool {
	di := sms.data[i]
	dj := sms.data[j]
	return di.k < dj.k && di.v < dj.k
}

func (sms *stringizeMapSlice) Swap(i, j int) {
	di := sms.data[i]
	sms.data[i] = sms.data[j]
	sms.data[j] = di
}

func (s *sha256Hasher) stringize(i interface{}) string {
	if i == nil {
		return "<nil>"
	}

	var ret string
	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		{
			l := []string{}
			sl := reflect.ValueOf(i)

			for i := 0; i < sl.Len(); i++ {
				l = append(l, s.stringize(sl.Index(i).Interface()))
			}
			sort.Strings(l)
			ret = fmt.Sprint(l)
		}
	case reflect.Map:
		{
			sms := &stringizeMapSlice{data: []*stringizeMapEntry{}}
			ml := reflect.ValueOf(i)

			for _, k := range ml.MapKeys() {
				v := ml.MapIndex(k)

				e := &stringizeMapEntry{
					k: s.stringize(k.Interface()),
					v: s.stringize(v.Interface()),
				}

				sms.data = append(sms.data, e)
			}
			sort.Sort(sms)
			ret = fmt.Sprint(sms.data)
		}
	default:
		ret = fmt.Sprint(i)
	}
	return ret
}

// Hash returns a Hash instance using the algorithm and salt
func (s *sha256Hasher) Hash(i interface{}) *Hash {
	h := sha256.New()
	h.Write(s.s)

	h.Write([]byte(s.stringize(i)))

	return &Hash{
		H: h.Sum(nil),
	}
}

type sha256HasherCreator struct {
}

// New returns a Hasher initialised with the specified salt
func (sc *sha256HasherCreator) New(salt []byte) Hasher {
	s := &sha256Hasher{}
	s.s = s.copyBytes(salt)
	return s
}

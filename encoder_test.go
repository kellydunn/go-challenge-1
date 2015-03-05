package drum

import ("testing"
       "path"
       "io/ioutil"
       "bytes"
       "os"
)

func TestEncodePattern(t *testing.T) {
	p := &Pattern{
		Version: "0.808-alpha",
		Tempo:   123.1,
	}

	err := EncodePattern(p, "test.splice")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDecodeEncode(t *testing.T){
	tData := []struct {
		path   string
		output string
	}{
                {"pattern_1.splice",
                 `Saved with HW Version: 0.808-alpha
Tempo: 120                      
(0) kick	|x---|x---|x---|x---|
(1) snare	|----|x---|----|x---|
(2) clap	|----|x-x-|----|----|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(4) hh-close	|x---|x---|----|x--x|
(5) cowbell	|----|----|--x-|----|
`,
		},
        }

        for _, exp := range tData {
            decoded, err := DecodeFile(path.Join("fixtures", exp.path))
            if err != nil {
               t.Errorf("Unexpected Error testing a full integration of decoding and encoding a splice file, %v", err)
            }

            err = EncodePattern(decoded, path.Join("fixtures", exp.path + "-encoded"))
            if err != nil {
               t.Errorf("Unexpected Error re-encoding a new file: %v", err)
            }

            fd, err := os.Open(path.Join("fixtures", exp.path + "-encoded"))
            if err != nil {
               
            }

            originalFixture, err := ioutil.ReadAll(fd)
            if err != nil {
               
            }
            
            newFd, err := os.Open(path.Join("fixtures", exp.path + "-encoded"))
            if err != nil {

            }

            newFixture, err := ioutil.ReadAll(newFd)
            if err != nil {

            }

            if bytes.Compare(originalFixture, newFixture) != 0 {
               t.Errorf("Old and new files differ")
            }
        }
}
package destroy

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//ModelCmd destroys a passed model
var ModelCmd = &cobra.Command{
	Use: "model [name]",
	//Example: "resource cars",
	Aliases: []string{"m"},
	Short:   "Destroys model files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you need to provide a valid model name in order to destroy it")
		}

		name := args[0]
		fileName := flect.Pluralize(flect.Underscore(name))

		removeModel(name)
		removeMigrations(fileName)

		return nil
	},
}

func confirm(msg string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')

	return (text == "y\n" || text == "Y\n")
}

func removeModel(name string) {
	if YesToAll || confirm("Want to remove model? (y/N)") {
		modelFileName := flect.Singularize(flect.Underscore(name))

		os.Remove(filepath.Join("models", fmt.Sprintf("%v.go", modelFileName)))
		os.Remove(filepath.Join("models", fmt.Sprintf("%v_test.go", modelFileName)))

		logrus.Infof("- Deleted %v\n", fmt.Sprintf("models/%v.go", modelFileName))
		logrus.Infof("- Deleted %v\n", fmt.Sprintf("models/%v_test.go", modelFileName))
	}
}

func removeMatch(folder, pattern string) {
	files, err := ioutil.ReadDir(folder)
	if err == nil {
		for _, f := range files {
			matches, _ := filepath.Match(pattern, f.Name())
			if !f.IsDir() && matches {
				path := filepath.Join(folder, f.Name())
				os.Remove(path)
				logrus.Infof("- Deleted %v\n", path)
			}
		}
	}
}

func removeMigrations(fileName string) {
	if YesToAll || confirm("Want to remove migrations? (y/N)") {
		removeMatch("migrations", fmt.Sprintf("*_create_%v.up.*", fileName))
		removeMatch("migrations", fmt.Sprintf("*_create_%v.down.*", fileName))
	}
}

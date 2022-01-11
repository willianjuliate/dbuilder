package newf

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type new_dir struct {
	name_project string
	qtd_modules  int
	dirs_default []string
}

func New() new_dir {
	nf := new_dir{
		name_project: "SEM_NOME",
		qtd_modules:  0,
		dirs_default: []string{
			"AUDIOS",
			"DOCUMENTOS",
			"IMAGENS",
			"MASCARAS",
			"PROJETOS FINAIS",
			"ROTEIROS",
			"VIDEOS",
			"ZIP",
		},
	}
	return nf
}
func (n *new_dir) NameProject(nomep string) string {
	if nomep != "" {
		n.name_project = nomep
		return fmt.Sprintf("NOME DO PROJETO: %s", n.name_project)
	}
	return fmt.Sprintf("NOME DO PROJETO: %s", n.name_project)
}
func (q *new_dir) QuantityModules(qtd string) string {
	num, _ := strconv.ParseInt(strings.TrimSpace(qtd), 10, 32)
	if num > 0 {
		q.qtd_modules = int(num)
		return fmt.Sprintf("VALOR DEFINIDO '%d'", num)
	} else {
		return "VALOR PADRÃO DEFINIDO 0"
	}

}
func (a *new_dir) AddNewDirs(names_dirs string) {
	if names_dirs != "" {
		for _, n := range strings.Split(names_dirs, ",") {
			dir := strings.TrimSpace(n)
			if dir != "" && dir != "," {
				a.dirs_default = append(a.dirs_default, dir)
			}
		}
		sort.Strings(a.dirs_default)
	}
}
func (r *new_dir) RemoveDirs(names_dirs string) {
	if names_dirs != "" {
		for _, n := range strings.Split(names_dirs, ",") {
			dir := strings.TrimSpace(n)
			i, name, err := r.getIndex(r.dirs_default, dir)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("# PASTA REMOVIDA '%s'\n", name)
				r.dirs_default = append(r.dirs_default[:i], r.dirs_default[i+1:]...)
			}
		}
	}
}
func (*new_dir) getIndex(s []string, name_dir string) (int, string, error) {
	for i, name := range s {
		if name_dir == name {
			return i, name, nil
		}
	}
	errof := fmt.Sprintf("# PASTA NÃO ENCONTRADO '%s' #", name_dir)
	return -1, name_dir, errors.New(errof)
}
func (s *new_dir) ShowDirs() {
	fmt.Printf("\n├─ %s\n", s.name_project)
	for _, name := range s.dirs_default {
		if name != "" {
			fmt.Printf("│     ├─ %s\n", name)
		}
	}
}
func (c *new_dir) CreateDirFiles() string {
	if existPath := os.Mkdir(c.name_project, os.ModePerm); existPath != nil {
		return fmt.Sprintf("\nPROJETO '%s' JÁ EXISTE!\n", c.name_project)
	} else {
		for _, root := range c.dirs_default {
			root_path := fmt.Sprintf("./%s/%s", c.name_project, root)
			os.Mkdir(root_path, os.ModePerm)
		}
		c.createFilesTscproj()
	}
	return fmt.Sprintf("# PROJETO '%s' CONSTRUÍDO #", c.name_project)
}
func (f *new_dir) createFilesTscproj() {
	str := `
{
	"title": "",
	"description": "",
	"author": "",
	"targetLoudness": -18.0,
	"shouldApplyLoudnessNormalization": true,
	"videoFormatFrameRate": 30,
	"audioFormatSampleRate": 44100,
	"width": 1440.0,
	"height": 900.0,
	"version": "5.0",
	"editRate": 30,
	"authoringClientName": {
		"name": "Camtasia",
		"platform": "Windows",
		"version": "21.0"
	},
	"timeline": {
		"id": 1,
		"sceneTrack": {
			"scenes": [
				{
					"csml": {
						"tracks": [
							{
								"trackIndex": 0,
								"medias": []
							},
							{
								"trackIndex": 1,
								"medias": []
							}
						]
					}
				}
			]
		},
		"captionAttributes": {
			"enabled": true,
			"fontName": "Arial",
			"fontSize": 53,
			"backgroundColor": [
				0,
				0,
				0,
				191
			],
			"foregroundColor": [
				255,
				255,
				255,
				255
			],
			"lang": "pt",
			"alignment": 0,
			"defaultFontSize": true,
			"opacity": 0.5,
			"backgroundEnabled": true,
			"backgroundOnlyAroundText": true
		},
		"gain": 1.0,
		"legacyAttenuateAudioMix": true,
		"backgroundColor": [
			0,
			0,
			0,
			255
		]
	},
	"metadata": {
		"AutoSaveFile": "",
		"CanvasZoom": 75,
		"Date": "",
		"IsAutoSave": false,
		"Language": "PTB",
		"ProfileName": "",
		"Title": "",
		"audioNarrationNotes": "",
		"calloutStyle": "Basic"
	}
}`
	for i := 1; i <= f.qtd_modules; i++ {
		root_name_files := fmt.Sprintf("./%s/%s_0%d.tscproj", f.name_project, f.name_project, i)
		os.WriteFile(root_name_files, []byte(str), os.ModePerm)
	}
}

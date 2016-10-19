package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

func UpdateISFile(path, name, tag string, latest bool) {
	log.Info("Updating IS file")
	var is ImageStream
	is.readISFile(path)
	is.createNewRecord(tag, name, latest)
	is.writeISFile(path)

}

func (is *ImageStream) readISFile(path string) {
	log.Infof("Read file %s", path)
	dat, err := ioutil.ReadFile(path)
	check(err)

	var imageStream ImageStream
	error := json.Unmarshal(dat, &imageStream)
	if error != nil {
		log.Fatal(error)
	}
	*is = imageStream
}

func (is *ImageStream) writeISFile(path string) {
	log.Infof("write file %s", path)
	data, err := json.MarshalIndent(is, "", "\t")
	if err != nil {
		log.Debug(err)

		return
	}
	err = ioutil.WriteFile(path, data, 0644)
	check(err)
}

func (is *ImageStream) createNewRecord(tagstr, name string, latest bool) {
	//check if we can find right image name if the file
	log.Debugf("Copy %s component for replacement", name)
	var item Items
	var tag Tags
	for _, v := range is.Items {
		if v.ItemsMetadata.Name == name {
			log.Debugf("Component %s found in List, copy it", name)
			item = v
		}
	}
	//error if image name was not found
	if len(item.ItemsMetadata.Name) == 0 {
		log.Fatalf("ImageStream %s not found in file provided", name)
	} else {
		log.Debug("Create new TAG to ammend existing list")
		//we need any tag with reference to dockerImage.
		for _, v := range item.Spec.Tags {
			if v.From.Kind == "DockerImage" {
				tag = item.Spec.Tags[0]

				//update copied element with new details
				if len(tag.Name) == 0 {
					log.Fatal("Error in reading copied tag")
				} else {
					log.Debugf("Overwrite details url:[%s] and name[%s] of the copy to %s", tag.From.Name, tag.Name, tagstr)
					tag.From.Name = strings.Replace(tag.From.Name, tag.Name, tagstr, 1)
					tag.Name = tagstr
					//append existing list with new element
					item.Spec.Tags = Extend(item.Spec.Tags, tag)
				}
			}
		}
		//if latest set to true we set latest to point to new date
		if latest {
			log.Infof("Latest set to repoint to new %s tag", tagstr)
			for k, v := range item.Spec.Tags {
				if v.Name == "latest" {
					item.Spec.Tags[k].From.Name = tagstr
					item.Spec.Tags[k].From.Kind = "ImageStreamTag"
				}
			}

			b, err := json.MarshalIndent(item, "", "\t")
			if err != nil {
				log.Debug(err)
				return
			}
			log.Debug(string(b))
		}

		//replace IS in default file with our new generated one
		for k, v := range is.Items {
			if v.ItemsMetadata.Name == name {
				log.Debugf("Replace %s with new generated item", name)
				is.Items[k] = item
			}
		}
	}
}

func Extend(slice []Tags, element Tags) []Tags {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]Tags, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func check(e error) {
	if e != nil {
		log.Fatalf("Error: %s", e)
		os.Exit(1)
	}
}

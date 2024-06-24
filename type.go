package main

import (
	"ParallelTar/job"
	"ParallelTar/tar"
	"log"
)

type Service struct {
	Dir  string
	Job  int
	Type string
}

func (svc *Service) Tar() {
	dirs, _ := listSubDirectories(svc.Dir)

	pool := job.NewPool(&svc.Job)
	for _, d := range dirs {
		if globalContinue[d] {
			continue
		}
		f := func(source, packType string) {
			err := tar.Tar(source, packType)
			if err != nil {
				log.Printf("%s, %s", source, err)
				record(".tar_error", source)
			}

			record(".continue", source)
			log.Printf("Packing %s completed.", source)
		}

		t := svc.Type
		pool.AddJobTar(f, d, t)
	}
	pool.Wait()
}

func (svc *Service) UnTar() {
	files, _ := listSubPackedFiles(svc.Dir)

	pool := job.NewPool(&svc.Job)
	for _, file := range files {
		if globalContinue[file] {
			continue
		}
		f := func(source string) {
			err := tar.UnTar(source)
			if err != nil {
				log.Printf("%s, %s", source, err)
				record(".untar_error", source)
			}

			record(".continue", source)
			log.Printf("Unpacking %s completed.", source)
		}

		pool.AddJobUnTar(f, file)
	}
	pool.Wait()
}

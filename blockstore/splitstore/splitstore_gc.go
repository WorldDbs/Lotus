package splitstore

import (
	"os"
	"path/filepath"
	"time"

	bstore "github.com/filecoin-project/lotus/blockstore"
)

func (s *SplitStore) gcHotstore() {
	// check if the hotstore is movable; if so, move it.
	if mover, ok := s.hot.(bstore.BlockstoreMover); ok {
		log.Info("moving hotstore")
		startMove := time.Now()
		err := mover.MoveTo("", nil)
		if err != nil {
			log.Warnf("error moving hotstore: %s", err)
			return
		}

		log.Infof("moving hotstore done", "took", time.Since(startMove))

		// clean up empty dirs in our path from previous compactions; MoveTo only removes the link
		log.Info("cleaning up splitstore directory")
		entries, err := os.ReadDir(s.path)
		if err != nil {
			log.Warnf("error reading splitstore directory (path: %s): %s", s.path, err)
			return
		}

		for _, e := range entries {
			path := filepath.Join(s.path, e.Name())

			if e.IsDir() && (e.Type()&os.ModeSymlink) == 0 {
				es, err := os.ReadDir(path)
				if err != nil {
					log.Warnf("error reading splitstore subdirectory %s: %s", e.Name(), err)
					continue
				}

				if len(es) > 0 {
					continue
				}

				err = os.Remove(path)
				if err != nil {
					log.Warnf("error removing splitstore subdirectory %s: %s", e.Name(), err)
				}

				log.Infof("removed empty splitstore subdirectory %s", e.Name())
			}
		}

		return
	}

	// check if the hotstore supports online GC; if so, GC it.
	if gc, ok := s.hot.(bstore.BlockstoreGC); ok {
		log.Info("garbage collecting hotstore")
		startGC := time.Now()
		err := gc.CollectGarbage()
		if err != nil {
			log.Warnf("error garbage collecting hotstore: %s", err)
			return
		}

		log.Infof("garbage collecting hotstore done", "took", time.Since(startGC))
		return
	}

	return
}

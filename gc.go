package main

import "time"

const (
	ProcessKeysCount = 100
	GarbageRate      = 0.25
)

func LazyGarbageCollect(db *Database, key string) {
	db.DelKeyWithExpire(key)
}

func StartCollectGarbageByTime(db *Database) {
	tk := time.NewTicker(100 * time.Millisecond)
	for {
		<-tk.C
		for {
			count := 0
			total := ProcessKeysCount

			DoProcess(func() {
				keys, num := db.getRandomExpireKeys(ProcessKeysCount)
				total = num

				for _, key := range keys {
					if db.CheckExpireValid(key) == KeyExpired {
						db.DelKeyWithExpire(key)
						count++
					}
				}
			})

			// if current garbage rate is high, do another round
			if float32(count)/float32(total) < GarbageRate {
				break
			}
		}
	}
}

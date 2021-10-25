package maker

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransNo(t *testing.T) {
	userId := 18922911
	transNo := TransNo(uint(userId))

	userIdStr := strconv.Itoa(userId)

	assert.Equal(t, len(transNo), 24)
	assert.Equal(t, transNo[0:2], fmt.Sprintf("%d", time.Now().Year()-2000))
	assert.Equal(t, transNo[2:5], fmt.Sprintf("%d", time.Now().YearDay()))
	assert.Equal(t, transNo[15:19], fmt.Sprintf(userIdStr[len(userIdStr)-4:]))
}

func TestTransNoWith(t *testing.T) {
	userId := 18922911

	transNo := TransNoWith(uint(userId), 16)
	assert.Equal(t, len(transNo), 16)

	userIdStr := fmt.Sprintf("%05d", userId)
	transNo2 := TransNoWith(uint(userId), 30)
	assert.Equal(t, len(transNo2), 30)
	assert.Equal(t, transNo2[0:2], fmt.Sprintf("%d", time.Now().Year()-2000))
	assert.Equal(t, transNo2[2:5], fmt.Sprintf("%d", time.Now().YearDay()))
	assert.Equal(t, transNo[len(transNo)-9:len(transNo)-5], fmt.Sprintf(userIdStr[len(userIdStr)-4:]))
}

func TestGenTransNoWith_Sync(t *testing.T) {
	for i := 16; i <= 32; i++ {
		_computeReplace(20, 1000, i, false)
	}
}

func _computeReplace(thread, max, transNoLen int, showReplace bool) {
	var wg sync.WaitGroup
	var set sync.Map

	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func(uid int) {
			defer wg.Done()

			for j := 0; j < max; j++ {
				transNo := TransNoWith(uint(uid), uint(transNoLen))
				v, ok := set.Load(transNo)
				if ok {
					set.Store(transNo, v.(int)+1)
				} else {
					set.Store(transNo, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	replaceCount := 0
	set.Range(func(key, value interface{}) bool {
		v, ok := value.(int)
		if !ok {
			return false
		}

		if v > 1 {
			replaceCount++
			if showReplace {
				log.Printf("%s \t %d\n", key, v)
			}
		}
		return true
	})

	total := thread * max
	rate := (float64(replaceCount) / float64(total)) * 100
	log.Println(fmt.Sprintf("len: %d, replace: %d, total: %d, repetitive rate: %0.6f%%", transNoLen, replaceCount, total, rate))
}

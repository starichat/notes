package pkg

import (
	"log"
	"testing"
)

func TestGenerateChannleId(t *testing.T) {
	log.Println(GenerateChannleId(1234,1,1,10000))
}

func TestGetType(t *testing.T) {
	ids := GenerateChannleId(1234,1,1,100);
	log.Println(ids)
	log.Println(GetType(ids))
	log.Println(GetConsumeAmount(ids))
}

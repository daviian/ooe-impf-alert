package main

import (
	"flag"
	"log"
	"sort"
	"time"
)

func main() {
	err := parseFlags()
	if err != nil {
		log.Println(err)
		flag.Usage()
		return
	}

	for range time.Tick(10 * time.Second) {
		checkForEarlierSlot()
	}
}

func checkForEarlierSlot() {
	slots := getSlotsForAuthorities()
	if len(slots) <= 0 {
		log.Println("no slots found")
		return
	}

	earliestSlot := getEarliestSlot(slots)

	if earliestSlot.StartDate.Before(date) {
		handleFoundSlot(earliestSlot)
		return
	}

	log.Printf("no earlier slot found. earliest slot on %s\n", earliestSlot.StartDate.Format(time.RFC822))
}

func getSlotsForAuthorities() []slot {
	var slots []slot
	for _, authority := range authorities {
		slotsOfAuthority, err := getCovidVaccinationSlots(authority)
		if err != nil {
			log.Fatal(err)
		}

		for _, slot := range slotsOfAuthority {
			slots = append(slots, slot)
		}
	}
	return slots
}

func getEarliestSlot(slots []slot) slot {
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].StartDate.Before(slots[j].StartDate)
	})

	return slots[0]
}

func handleFoundSlot(earliestSlot slot) {
	log.Printf("YEAHHH. found an earlier slot on %s\n", earliestSlot.StartDate.Format(time.RFC822))

	err := sendPushNotification(earliestSlot.StartDate)
	if err != nil {
		log.Fatal(err)
	}
}

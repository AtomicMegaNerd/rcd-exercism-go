package encode

import (
	"fmt"
	"regexp"
	"strconv"
)

type Cluster struct {
	ch    rune
	count int
}

func NewCluster(ch rune) Cluster {
	return Cluster{ch: ch, count: 1}
}

func NewClusterWithCount(ch rune, count int) Cluster {
	return Cluster{ch: ch, count: count}
}

func (c *Cluster) increment() {
	c.count++
}

func (c *Cluster) ToEncodedString() string {
	if c.count == 1 {
		return string(c.ch)
	}
	return fmt.Sprintf("%d%c", c.count, c.ch)
}

func (c *Cluster) ToDecodedString() string {
	decodedString := ""
	for ix := 0; ix < c.count; ix++ {
		decodedString += string(c.ch)
	}
	return decodedString
}

func RunLengthEncode(input string) string {
	encodedString := ""
	clusters := make([]Cluster, 0)
	var currentCluster Cluster

	for ix, ch := range input {
		if ix == 0 {
			currentCluster = NewCluster(ch)
			continue
		}

		if ch != currentCluster.ch {
			clusters = append(clusters, currentCluster)
			currentCluster = NewCluster(ch)
		} else {
			currentCluster.increment()
		}

		// Don't forget about the last rune in the string!
		if ix == len(input)-1 {
			clusters = append(clusters, currentCluster)
		}
	}

	for _, cluster := range clusters {
		encodedString += cluster.ToEncodedString()
	}
	return encodedString
}

func RunLengthDecode(input string) string {
	decodedString := ""
	clusters := make([]Cluster, 0)
	var currentCluster Cluster

	pattern := regexp.MustCompile("(\\d+)([a-zA-z]|\\s+)|([a-zA-z]|\\s+)")

	matches := pattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		// If the first match is a number we want to capture that
		count, err := strconv.Atoi(match[1])
		if err == nil {
			ch := []rune(match[2])[0]
			currentCluster = NewClusterWithCount(ch, count)
		} else {
			ch := []rune(match[0])[0]
			currentCluster = NewClusterWithCount(ch, 1)
		}
		clusters = append(clusters, currentCluster)
	}

	for _, cluster := range clusters {
		decodedString += cluster.ToDecodedString()
	}
	return decodedString
}

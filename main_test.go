package main

import (
    "testing"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

func TestSpace(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Space Suite")
}

var _ = Describe("Space", func() {

})

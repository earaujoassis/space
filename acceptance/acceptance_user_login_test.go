package acceptance

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    . "github.com/sclevine/agouti/matchers"
    "github.com/sclevine/agouti"
)

var _ = Describe("UserLogin", func() {
    var page *agouti.Page

    BeforeEach(func() {
        var err error
        page, err = agoutiDriver.NewPage()
        Expect(err).NotTo(HaveOccurred())
    })

    AfterEach(func() {
        Expect(page.Destroy()).To(Succeed())
    })

    It("should manage user authentication", func() {
        By("redirecting the user to the login form from the webapp page", func() {
            Expect(page.Navigate(AcceptanceSettings("webappUrl"))).To(Succeed())
        })

        By("allowing the user to fill out the login form and submit it", func() {
            Eventually(page.FindByName("email")).Should(BeFound())
            Expect(page.FindByName("email").Fill(AcceptanceSettings("email"))).To(Succeed())
            Expect(page.FindByName("password").Fill(AcceptanceSettings("password"))).To(Succeed())
            Expect(page.FirstByButton("Get Access").Submit()).To(Succeed())
        })

        By("allowing the user to view its data", func() {
            Eventually(page.Find("#user-context")).Should(BeFound())
            Expect(page.Find("#user-context").Click()).To(Succeed())
            Expect(page.Find(".user-submenu")).To(BeVisible())
            Expect(page.Find(".user-submenu .submenu-data p").Text()).To(Equal("Signed in as\nAcceptance Tests"))
        })

        By("allowing the user to logout", func() {
            Eventually(page.Find("#user-context")).Should(BeFound())
            Expect(page.Find("#user-context").Click()).To(Succeed())
            Expect(page.Find(".user-submenu")).To(BeVisible())
            Eventually(page.Find(".user-submenu ul.submenu a[href='/logout']")).Should(BeFound())
            Expect(page.Find(".user-submenu ul.submenu a[href='/logout']").Click()).To(Succeed())
            Eventually(page.FindByName("email")).Should(BeFound())
            Eventually(page.FindByName("password")).Should(BeFound())
            Eventually(page.FirstByButton("Get Access")).Should(BeFound())
        })
    })
})

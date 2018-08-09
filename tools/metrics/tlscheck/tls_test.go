package tlscheck_test

import (
	"crypto/tls"

	. "github.com/alphagov/paas-cf/tools/metrics/tlscheck"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TLSCheck", func() {
	var (
		checker *TLSChecker
	)

	BeforeEach(func() {
		checker = &TLSChecker{}
	})

	Context("DaysUntilExpiry", func() {
		It("returns >0 for non-expired certificate", func() {
			daysUntilExpiry, err := checker.DaysUntilExpiry("badssl.com:443", &tls.Config{})
			Expect(err).NotTo(HaveOccurred())
			Expect(daysUntilExpiry).To(BeNumerically(">", float64(0)))
		})

		It("returns 0 for expired certificate", func() {
			daysUntilExpiry, err := checker.DaysUntilExpiry("expired.badssl.com:443", &tls.Config{})
			Expect(err).NotTo(HaveOccurred())
			Expect(daysUntilExpiry).To(Equal(float64(0)))
		})

		It("returns error for certificate with incorrect common name", func() {
			_, err := checker.DaysUntilExpiry("wrong.host.badssl.com:443", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})

		It("returns error for certificate with untrusted root CA", func() {
			_, err := checker.DaysUntilExpiry("untrusted-root.badssl.com:443", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})

		// FIXME: This should be reverted back, once the https://github.com/chromium/badssl.com/issues/359
		// has been resolved.
		//
		//		It("returns error for certificate with self-signed CA", func() {
		//			_, err := checker.DaysUntilExpiry("self-signed.badssl.com:443", &tls.Config{})
		//			Expect(err).To(HaveOccurred())
		//		})

		It("returns error for certificate with null cipher suite", func() {
			_, err := checker.DaysUntilExpiry("null.badssl.com:443", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})

		It("returns err when cannot connect", func() {
			_, err := checker.DaysUntilExpiry("no.connection.invalid:443", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})

		It("returns err when addr is a URL", func() {
			_, err := checker.DaysUntilExpiry("http://badssl.com:443", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})

		It("returns err when addr has no port", func() {
			_, err := checker.DaysUntilExpiry("badssl.com", &tls.Config{})
			Expect(err).To(HaveOccurred())
		})
	})
})

package configs

import "crypto/tls"

func GetTlsConfigWithCer(cer tls.Certificate) *tls.Config {
	return &tls.Config{
		MinVersion: tlsVersion(),
		CurvePreferences:  curvePreferences(),
		PreferServerCipherSuites: cipherSuitesPreference(),
		CipherSuites: cipherSuites(),
		Certificates: []tls.Certificate{cer},
	}
}

func GetTlsConfigNoCer() *tls.Config {
	return &tls.Config{
		MinVersion:	tlsVersion(),
		CurvePreferences: curvePreferences(),
		PreferServerCipherSuites: 	cipherSuitesPreference(),
		CipherSuites: cipherSuites(),
	}
}

func cipherSuitesPreference() bool {
	return true
}

func tlsVersion() uint16{
	return tls.VersionTLS12
}

func cipherSuites() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	}
}

func curvePreferences() []tls.CurveID {
	return []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}
}

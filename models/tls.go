package models

// TlsConfig - TLS configuration settings and status
type TlsConfig struct {
	Enabled           bool     `json:"enabled" description:"enabled is the encryption (DoT/DoH/HTTPS) status"`
	ServerName        string   `json:"server_name" description:"server_name is the hostname of your HTTPS/TLS server"`
	ForceHttps        bool     `json:"force_https" description:"if true, forces HTTP->HTTPS redirect"`
	PortHttps         uint16   `json:"port_https" description:"HTTPS port. If 0, HTTPS will be disabled."`
	PortDnsOverTls    uint16   `json:"port_dns_over_tls" description:"DNS-over-TLS port. If 0, DoT will be disabled."`
	PortDnsOverQuic   uint16   `json:"port_dns_over_quic" description:"DNS-over-QUIC port. If 0, DoQ will be disabled."`
	CertificateChain  string   `json:"certificate_chain" description:"Base64 string with PEM-encoded certificates chain"`
	PrivateKey        string   `json:"private_key" description:"Base64 string with PEM-encoded private key"`
	PrivateKeySaved   bool     `json:"private_key_saved" description:"Set to true if the user has previously saved a private key as a string. This is used so that the server and the client don't have to send the private key between each other every time, which might lead to security issues."`
	CertificatePath   string   `json:"certificate_path" description:"Path to certificate file"`
	PrivateKeyPath    string   `json:"private_key_path" description:"Path to private key file"`
	ValidCert         bool     `json:"valid_cert" description:"Set to true if the specified certificates chain is a valid chain of X509 certificates."`
	ValidChain        bool     `json:"valid_chain" description:"Set to true if the specified certificates chain is verified and issued by a known CA."`
	Subject           string   `json:"subject" description:"The subject of the first certificate in the chain."`
	Issuer            string   `json:"issuer" description:"The issuer of the first certificate in the chain."`
	NotBefore         string   `json:"not_before" description:"The NotBefore field of the first certificate in the chain."`
	NotAfter          string   `json:"not_after" description:"The NotAfter field of the first certificate in the chain."`
	DnsNames          []string `json:"dns_names" description:"The value of SubjectAltNames field of the first certificate in the chain."`
	ValidKey          bool     `json:"valid_key" description:"Set to true if the key is a valid private key."`
	KeyType           string   `json:"key_type" description:"Key type. Can be one of: RSA, ECDSA."`
	WarningValidation string   `json:"warning_validation" description:"A validation warning message with the issue description."`
	ValidPair         bool     `json:"valid_pair" description:"Set to true if both certificate and private key are correct."`
	ServePlainDns     bool     `json:"serve_plain_dns" description:"Set to true if plain DNS is allowed for incoming requests."`
}

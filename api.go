package twinfield

import (
	"fmt"
	"time"

	"github.com/tim-online/go-twinfield/soap"
)

const sessionTimeout = time.Minute * 30

var (
	sessionID    string
	sessionStart time.Time
	User         = ""
	Password     = ""
	Organisation = ""
	cluster      = ""
)

func Login() error {
	// @TODO: implement retry strategy?
	session, err := login()
	sessionID = session
	return err
}

// @TODO: add KeepAlive() in separate thread or shutdown session after
// # of connections = 0
func login() (string, error) {
	session := soap.NewSessionSoap("", true, nil)
	logon := &soap.Logon{
		User:         User,
		Password:     Password,
		Organisation: Organisation,
	}

	logonResp, err := session.Logon(logon)
	if err != nil {
		return "", err
	}

	if *logonResp.LogonResult != soap.LogonResultOk {
		return "", fmt.Errorf("logonResult: %s", *logonResp.LogonResult)
	}
	cluster = logonResp.Cluster
	sessionID := logonResp.Envelope.Header.Header.(*soap.Header).SessionID
	return sessionID, nil
}

func GetSessionID() (string, error) {
	if sessionID == "" {
		err := Login()
		if err == nil {
			// logon succesful: update session start
			sessionStart = time.Now().UTC()
		}
		return sessionID, err
	}

	sessionAge := time.Now().UTC().Sub(sessionStart)
	if sessionAge > sessionTimeout {
		err := Login()
		if err == nil {
			// logon succesful: update session start
			sessionStart = time.Now().UTC()
		}
		return sessionID, err
	}

	return sessionID, nil
}

func GetOffices() ([]soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: "<list><type>offices</type></list>",
	}

	process, err := GetProcessXmlSoap()
	if err != nil {
		return nil, err
	}

	processResp, err := process.ProcessXmlString(xml)
	if err != nil {
		return nil, err
	}

	data := []byte(processResp.ProcessXmlStringResult)
	offices, err := soap.OfficesFromXml(data)
	return offices, err
}

func GetProcessXmlSoap() (*soap.ProcessXmlSoap, error) {
	sessionID, err := GetSessionID()
	if err != nil {
		return nil, err
	}
	url := cluster + "/webservices/processxml.asmx"
	process := soap.NewProcessXmlSoap(url, true, nil)
	header := &soap.Header{
		SessionID: sessionID,
	}
	process.SetHeader(header)
	return process, nil
}

func GetFinderXmlSoap() (*soap.FinderSoap, error) {
	sessionID, err := GetSessionID()
	if err != nil {
		return nil, err
	}
	url := cluster + "/webservices/finder.asmx"
	finder := soap.NewFinderSoap(url, true, nil)
	header := &soap.Header{
		SessionID: sessionID,
	}
	finder.SetHeader(header)
	return finder, nil
}

func GetDeclarationsXmlSoap() (*soap.DeclarationsSoap, error) {
	sessionID, err := GetSessionID()
	if err != nil {
		return nil, err
	}
	url := cluster + "/webservices/declarations.asmx"
	req := soap.NewDeclarationsSoap(url, true, nil)
	header := &soap.Header{
		SessionID: sessionID,
	}
	req.SetHeader(header)
	return req, nil
}

func GetOfficeByID(officeID string) (*soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: fmt.Sprintf(`<read>
			<type>office</type>
			<code>%s</code>
		</read>`, officeID),
	}

	process, err := GetProcessXmlSoap()
	if err != nil {
		return nil, err
	}

	processResp, err := process.ProcessXmlString(xml)
	if err != nil {
		return nil, err
	}

	data := []byte(processResp.ProcessXmlStringResult)
	office, err := soap.OfficeFromXml(data)
	return office, err
}

package aws

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/route53"
    "testing"
)

func TestDNS(t *testing.T) {
    sess, err := session.NewSession(
        &aws.Config{
            CredentialsChainVerboseErrors: aws.Bool(true),
        })
    if err != nil {
        panic(err)
    }

    svc := route53.New(sess)
    req := &route53.ChangeResourceRecordSetsInput{}
    req.SetHostedZoneId("Z10212951OIAJAFXITIEX")
    changeBatch := &route53.ChangeBatch{}
    changeBatch.SetComment("optional comment about the changes in this change batch request")
    changes := make([]*route53.Change, 0)
    change := &route53.Change{}
    change.SetAction("UPSERT")
    resourceRecordSet := &route53.ResourceRecordSet{}
    resourceRecordSet.SetName("_dnslink.dev.phexsbtmd.com")
    resourceRecordSet.SetType("TXT")
    resourceRecordSet.SetTTL(30)
    resourceRecord := &route53.ResourceRecord{}
    resourceRecord.SetValue("\"dnslink=/ipfs/QmRLLtevbC7VsPCriRzr3HPQWR5NkBxJxQFswxAn5vN4GT\"")
    resourceRecords := make([]*route53.ResourceRecord,0)
    resourceRecords = append(resourceRecords, resourceRecord)
    resourceRecordSet.SetResourceRecords(resourceRecords)
    change.SetResourceRecordSet(resourceRecordSet)
    changes = append(changes, change)
    changeBatch.SetChanges(changes)
    req.SetChangeBatch(changeBatch)
    rsp, err := svc.ChangeResourceRecordSets(req)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", rsp.String())
}

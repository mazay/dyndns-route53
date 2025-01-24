package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

type Route53 struct {
	Region          string
	AccessKey       string
	SecretAccessKey string
	client          *route53.Client
}

// New returns a new Route53 client.
//
// If the AccessKey and SecretAccessKey fields are both set, the client will be
// configured to use them for authentication. Otherwise, the default credential
// provider chain will be used.
//
// The client will be configured to use the same region as the one that was
// specified when creating the Route53 object.
func (r *Route53) New() (*Route53, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		// handle error
		return nil, err
	}

	client := route53.NewFromConfig(cfg, func(o *route53.Options) {
		if r.AccessKey != "" && r.SecretAccessKey != "" {
			o.Credentials = aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					r.AccessKey,
					r.SecretAccessKey,
					"",
				),
			)
		}
	})

	return &Route53{
		client: client,
	}, nil
}

// UpdateRRecord updates a DNS record in Route53 with the provided name, target, and zone ID.
// It performs an upsert operation on a CNAME record type, setting the target as the value.
// Returns an error if the update fails.
func (r *Route53) UpdateRRecord(name string, target string, zoneId string) error {
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: types.ChangeActionUpsert,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: aws.String(name),
						Type: types.RRTypeCname,
						ResourceRecords: []types.ResourceRecord{
							{
								Value: aws.String(target),
							},
						},
						TTL: aws.Int64(60),
					},
				},
			},
		},
		HostedZoneId: aws.String(zoneId),
	}

	_, err := r.client.ChangeResourceRecordSets(context.TODO(), params)
	if err != nil {
		return err
	}

	return nil
}

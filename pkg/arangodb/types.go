package arangodb

import (
	"github.com/sbezverk/gobmp/pkg/base"
	"github.com/sbezverk/gobmp/pkg/bgp"
)

type peerToObject struct {
	Key         string              `json:"_key"`
	From        string              `json:"_from"`
	To          string              `json:"_to"`
	LocalBGPID  string              `json:"local_bgp_id"`
	RemoteBGPID string              `json:"remote_bgp_id"`
	LocalIP     string              `json:"local_ip"`
	RemoteIP    string              `json:"remote_ip"`
	BaseAttrs   *bgp.BaseAttributes `json:"base_attrs"`
	LocalASN    uint32              `json:"local_asn"`
	RemoteASN   uint32              `json:"reote_asn"`
	OriginAS    int32               `json:"origin_as"`
	ProtocolID  base.ProtoID        `json:"protocol_id"`
	Nexthop     string              `json:"nexthop"`
	Labels      []uint32            `json:"labels"`
	Name        string              `json:"name"`
	Session     string              `json:"session"`
}

type peerFromObject struct {
	Key         string              `json:"_key"`
	From        string              `json:"_from"`
	To          string              `json:"_to"`
	LocalBGPID  string              `json:"local_bgp_id"`
	RemoteBGPID string              `json:"remote_bgp_id"`
	LocalIP     string              `json:"local_ip"`
	RemoteIP    string              `json:"remote_ip"`
	BaseAttrs   *bgp.BaseAttributes `json:"base_attrs"`
	LocalASN    uint32              `json:"local_asn"`
	RemoteASN   uint32              `json:"reote_asn"`
	OriginAS    int32               `json:"origin_as"`
	ProtocolID  base.ProtoID        `json:"protocol_id"`
	Nexthop     string              `json:"nexthop"`
	Labels      []uint32            `json:"labels"`
	Name        string              `json:"name"`
	Session     string              `json:"session"`
}

type unicastPrefixEdgeObject struct {
	Key        string              `json:"_key"`
	From       string              `json:"_from"`
	To         string              `json:"_to"`
	Prefix     string              `json:"prefix"`
	PrefixLen  int32               `json:"prefix_len"`
	LocalIP    string              `json:"router_ip"`
	PeerIP     string              `json:"peer_ip"`
	BaseAttrs  *bgp.BaseAttributes `json:"base_attrs"`
	PeerASN    uint32              `json:"peer_asn"`
	OriginAS   int32               `json:"origin_as"`
	ProtocolID base.ProtoID        `json:"protocol_id"`
	Nexthop    string              `json:"nexthop"`
	Labels     []uint32            `json:"labels"`
	Name       string              `json:"name"`
	PeerName   string              `json:"peer_name"`
}

type inetPrefix struct {
	ID        string `json:"_id,omitempty"`
	Key       string `json:"_key,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	PrefixLen int32  `json:"prefix_len,omitempty"`
	OriginAS  int32  `json:"origin_as"`
}

type inetPrefixEdgeObject struct {
	Key        string              `json:"_key"`
	From       string              `json:"_from"`
	To         string              `json:"_to"`
	Prefix     string              `json:"prefix"`
	PrefixLen  int32               `json:"prefix_len"`
	LocalIP    string              `json:"router_ip"`
	PeerIP     string              `json:"peer_ip"`
	BaseAttrs  *bgp.BaseAttributes `json:"base_attrs"`
	PeerASN    uint32              `json:"peer_asn"`
	OriginAS   int32               `json:"origin_as"`
	ProtocolID base.ProtoID        `json:"protocol_id"`
	Nexthop    string              `json:"nexthop"`
	Labels     []uint32            `json:"labels"`
	Name       string              `json:"name"`
	PeerName   string              `json:"peer_name"`
}

type bgpPeer struct {
	Key             string         `json:"_key,omitempty"`
	ID              string         `json:"_id,omitempty"`
	BGPRouterID     string         `json:"router_id,omitempty"`
	ASN             int32          `json:"asn"`
	AdvCapabilities bgp.Capability `json:"adv_cap,omitempty"`
}

type bgpPrefix struct {
	ID        string `json:"_id,omitempty"`
	Key       string `json:"_key"`
	Prefix    string `json:"prefix"`
	PrefixLen int32  `json:"prefix_len"`
	OriginAS  int32  `json:"origin_as"`
	RouterID  string `json:"router_id"`
}

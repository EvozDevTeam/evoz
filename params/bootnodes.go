// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main network.
var MainnetBootnodes = []string{
	"enode://7b36f314918a282156bc9bbb8c2decb45fee7618829debd06adf2cbb17a355a3341fbcccf8e85e780e47eb4faaca737199573eb2e8ad11eab33660644299e523@109.123.237.179:30303",
	"enode://268441dcfeb7672f450cdb8be7fddc3189b1b674f9bddb37b8d063a0815abb688851ee447e76a414ff4ee12f78c0bba001508587137d045d5afbc5f5f36cdf0c@46.250.227.45:30303",
	"enode://ddab5b652bc9904f04d29ecd83014237ac647899c1c1a4a12ec2564d3fab52db8cddff5860c1bdaf3ac0f713b1d77e056ab0a0569b9cf927037c1ddf40dea3f3@84.247.146.160:30303",
	"enode://cd33984d7c4f977a8259fc62f653b70538df3a23857f274f7eb68135c4c44bc3f5feb5f454fb73e713077cacb81a1b07f1676d470956ce9f5c7f028157ed8b4f@84.247.148.103:30303",
	"enode://d3320558d40acaecd6951ce60c35dfcd798c4e56f6c7ae4cfb64772aab8ee9dd66e6bf2cd63b0b55b967bbf47e5cda26028eea9d1d65b8f9dc16ad4fe8757940@84.247.149.227:30303",
	"enode://5b12c8f7fb7f8b38316cda3981fbf5e8ace9a5aab647b36aee1d8dfe917ad9eec0bb556005d653e6d2481d6ac6135ba728a222d269f8c49fbe544e75e6c17b75@84.247.149.229:30303",
	"enode://94efa794852371382c91574772dd58b3360bc9450654555d20683daf8f27e952543cb721b45dea1d3e6551e477c8a6550e1723fc47300585f9f227d9c36064b8@217.15.161.70:30303",
	"enode://e23a9cb36aca4778ff2158586cf3d1176a3d93956329fab9e96993c333f4b6b48bc2d2d95a6c469ef5d3832139697747de112fc268b743a43558b8748a7c87a5@217.15.161.132:30303",
	"enode://740dd3110b79844cb0a31bc5da119abddcbdd5362e80a8fbe67763a429c431ba85c9fd9c28cfaad22bd2edc08d92bd56de4fc5c6d05c677cadc0dcf61a2c7136@217.15.161.214:30303",
	"enode://0a76aecd1df0d524de7db92c7a14abd05cbcf3e2163d8e54b0deb62e0db92aaadaf326b26a545bb8558b09b523fb262809f94cb01123e712dfe133805cba524c@217.15.162.116:30303",
	"enode://b1911758273339952f661af397afabe0e090381b72f9743f63592dda70edb0f4d24fb448614e33f471b83fa6eb5814b6a073bbdf9672b4ee9bfa6d3b1fb3cb07@217.15.162.241:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}

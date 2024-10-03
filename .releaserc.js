module.exports = {
    "branches": [
        {
            "name": "develop",
            "prerelease": true
        },
        {
            "name": "main",
            "prerelease": false
        }
    ],
    "plugins": [
        "@semantic-release/commit-analyzer",
        "@semantic-release/release-notes-generator",
        [
            "@bpgeck/semantic-release-kaniko",
            {
                "destination": [
                    "sammobach/gitdesk-forwarder:${version}",
                    "sammobach/gitdesk-forwarder:latest"
                ]
            }
        ],
        [
            "@semantic-release/git",
            {
                "assets": [
                    "CHANGELOG.md"
                ]
            }
        ],
        "@semantic-release/github"
    ],
    "repositoryUrl": "https://github.com/nilpntr/gitdesk-forwarder.git",
}
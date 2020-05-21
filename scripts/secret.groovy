openshift.withCluster( 'mycluster' ) {

        // You can use sub-verbs of create for some simple objects
        openshift.create( 'serviceaccount', 'my-service-account5' )

        // But you want to craft your own, don't you? First,
        // model it with Groovy Maps, Lists, and primitives.
        def secret = [
            "apiVersion": "v1",
            "kind": "Secret",
            "metadata": [
                "name": "mysecret5",
                "labels": [
                    'findme':'please'
                ]
            ],
            "stringData": [
                "secretKey": "YWRtaW4K",
                "secret": "redhat"
            ],
            "type": "Opaque"
        ]

        // create will marshal the model into JSON and send it to the API server.
        // We will add some passthrough arguments (--save-config and --validate)
        // just for fun.
        def objs = openshift.create( secret,'--save-config', '--validate' )

        // create(...) returns a Selector will select the resulting object(s).
        objs.describe()

        // But, you say, I've already modeled my object in JSON/YAML! It is in
        // an SCM or accessible with via HTTP, or..., or ...
        // Don't worry. Just get it to the current Jenkins workspace any way
        // you want (e.g. using a Jenkins plugin for your SCM). Then read the
        // file into a String using normal Jenkins steps.

        def fromJSON = openshift.create( '-f','https://raw.githubusercontent.com/jitendar-singh/jenkins-test/jenkins-functionality/templates/mysecret.yaml' )


}

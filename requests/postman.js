var path = require('path'),
    async = require('async'), //https://www.npmjs.com/package/async
    newman = require('newman'),

    parametersForTestRun = {
        collection: path.join(__dirname, 'postman_collection.json'), // your collection
        environment: path.join(__dirname, 'postman_environment.json'), //your env
    };

parallelCollectionRun = function(done) {
    newman.run(parametersForTestRun, done);
};

// Runs the Postman sample collection thrice, in parallel.
async.parallel([
        parallelCollectionRun,
        parallelCollectionRun,
        parallelCollectionRun
    ],
    function(err, results) {
        err && console.error(err);

        results.forEach(function(result) {
            var failures = result.run.failures;
            console.info(failures.length ? JSON.stringify(failures.failures, null, 2) :
                `${result.collection.name} ran successfully.`);
        });
    });
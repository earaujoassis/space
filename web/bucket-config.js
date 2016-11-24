var bucketAccess = process.env.SPACE_BUCKET_ACCESS;
var keySecret;

if (bucketAccess && bucketAccess.length) {
    keySecret = bucketAccess.split(':');
} else {
    keySecret = [null, null];
}

module.exports = {
    accessKeyId: keySecret[0],
    secretAccessKey: keySecret[1]
}

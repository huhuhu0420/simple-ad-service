import http from 'k6/http';

export const options = {
    scenarios: {
        my_scenario1: {
            executor: 'constant-arrival-rate',
            duration: '13s', // total duration
            preAllocatedVUs: 250, // to allocate runtime resources  
            rate: 11000, // number of constant iterations given `timeUnit`
            timeUnit: '1s',
        },
    },
};

let url = 'http://localhost:5000/api/v1/ad?Limit=3&Offset=0&Gender=M&Platform=iOS&Country=TW';

export default function () {
    http.get(url)
}

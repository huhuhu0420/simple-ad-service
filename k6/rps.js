import { check } from 'k6';
import http from 'k6/http';

export const options = {
    scenarios: {
        my_scenario1: {
            executor: 'constant-arrival-rate',
            duration: '30s', // total duration
            preAllocatedVUs: 100, // to allocate runtime resources     preAll

            rate: 10000, // number of constant iterations given `timeUnit`
            timeUnit: '1s',
        },
    },
};

export default function () {
    let response = http.get('http://localhost:5000/api/v1/ad');
    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}

import { sleep } from 'k6';
import http from 'k6/http';

export const options = {
    scenarios: {
        my_scenario1: {
            executor: 'constant-arrival-rate',
            duration: '10s', // total duration
            preAllocatedVUs: 300, // to allocate runtime resources     preAll
            rate: 12000, // number of constant iterations given `timeUnit`
            timeUnit: '1s',
        },
    },
};

let url = 'http://34.49.138.159/api/v1/ad';

export default function () {
    http.get(url)
}

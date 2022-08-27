import http from "k6/http";
import { check } from "k6";

export default function () {
  check(http.put("http://192.168.0.102:8082/mockserver/status"), {
    "status is 200": (r) => r.status == 200,
  });
}

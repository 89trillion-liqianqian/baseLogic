from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)



    @task
    def keyword(self):
        self.client.get("/getIntByStr?str=5/2/2")
---
title: Service Downtime Due to Insufficient Filesystem Capacity
slug: service-outage-due-to-storage
image: /assets/images/posts/post-316/210195173-9bc2975f-73f4-495d-bb19-71732ce593f2.png
date: 2023-05-07T00:00:00Z
author: bmf-san
categories:
  - Incident Report
tags:
  - Postmortem
translation_key: service-outage-due-to-storage
---

# Status
Resolved

# Incident
On January 2, 2023, around noon, it was noticed that accessing `https://bmf-tech.com/` resulted in slow responses and constant 500 errors. Upon attempting to log in to Grafana for investigation, login was unsuccessful.

Considering the possibility that some containers might have gone down for some reason, a deployment was carried out, but upon checking the error log, `no space left on device` was found, leading to the assumption that it was due to another cause.

# Impact
All services of bmf-tech became unavailable.

Checking the Nginx request status, it was confirmed that the service had been down since around 5:48 AM on January 2, 2023.

Restoration occurred around 12:40 PM on the same day.

![Screenshot 2023-01-02 13 26 53](https://user-images.githubusercontent.com/13291041/210195174-bf6c78c5-505d-41ff-8329-7781ef1fcae1.png)

Between 5:48 AM and 12:40 PM on January 2, 2023, 58 instances of 500 errors occurred.
*Note: I wanted to measure a certain number of users, but it was difficult to investigate as logs and GA4 were not adjusted for aggregation.*

![Screenshot 2023-01-02 13 48 34](https://user-images.githubusercontent.com/13291041/210195945-c9e1ae53-d624-4119-b8c5-22a185e66239.png)

# Cause
![Screenshot 2023-01-02 12 39 03](https://user-images.githubusercontent.com/13291041/210195173-9bc2975f-73f4-495d-bb19-71732ce593f2.png)

The cause was a lack of free space in the filesystem.

# Response
Deleted data that was occupying filesystem capacity to secure free space.

# Actions
1. SSH into the server and check disk free space using the df command. There was no free space.
   
   ```
   df -h
   ```
2. Investigated the areas likely using the most disk space with the du command. It was identified to be under `/var/lib/docker/`.
   
   ```
   du -sh /var/lib/docker*
   ```
3. Deleted unused Docker objects to secure free space.
   
   ```sh
   docker system prune -a
   ```
4. Additionally, deleted journal logs that were consuming the second most space, leaving 200M.
   
   ```sh
   journalctl --vacuum-size=200M
   ```

# Addendum
The above response was insufficient.

Logs accumulating under `/var/lib/docker/containers` were taking up significant space, so adjustments were made to the container log rotation.
   
   ```yml
   ex. 
     app:
       container_name: "app"
       logging:
         driver: "json-file"
         options:
           tag: "{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}"
           max-size: "10m" // File size to roll over
           max-file: "3" // Number of times logs can roll over before being discarded
   ```

This adjustment allowed for a significant increase in available capacity. This was the main bottleneck...

# Prevention
Added filesystem usage rate to alerts to detect and address issues in advance.

![Screenshot 2023-01-02 14 10 54](https://user-images.githubusercontent.com/13291041/210196950-64e37302-38a2-4cda-a34e-e20b99b77791.png)

# Others
I want to identify data that can be deleted and data that should be rotated to save space.

# Thoughts
This was the first incident since replacing bmf-tech.
# Introduction

This repository is a basic setup of how three servers (in this instance three local docker containers running Debian) would deploy a node for a queue, a node for a consumer and a node for a publisher.

This simple setup assumes that the server is running the same x86 architecture as the manager node.

The consumer and publisher have bene written in go as this is the language I am currently trying to get more familiar with.

Simple Debian docker containers were used as this project aims to replicate a deployment to actual servers which often run Debian or Ubuntu (or any other popular Linux distribution) and do not contain all the other necessities that are contained in docker images.

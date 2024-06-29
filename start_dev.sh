#!/bin/bash

(
	cd client || exit
	pnpm dev &
)

air

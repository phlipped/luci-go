#!/bin/bash
# Copyright 2016 The LUCI Authors. All rights reserved.
# Use of this source code is governed under the Apache License, Version 2.0
# that can be found in the LICENSE file.

TESTS=common/dirtools/tests/*.json

echo "Generating the test directories"
TMPDIR=/usr/local/google/tmp/luci-tests
mkdir -p $TMPDIR
for TESTFILE in $TESTS; do
	TESTNAME="$(basename $TESTFILE .json)"
	TESTDIR="$TMPDIR/$TESTNAME"
	if ! [ -d $TESTDIR ]; then
		echo "Generating test directory for $TESTNAME"
		gendir -config $TESTFILE -outdir $TESTDIR
	fi
done

for METHOD in simple nostat parallel; do
	echo "Running $METHOD"
	for TESTFILE in $TESTS; do
		TESTNAME="$(basename $TEST .json)"
		TESTDIR="$TMPDIR/$TESTNAME"
		OUTPUT=output.$METHOD.$TESTNAME
		$(which time) --verbose --output=$OUTPUT --append walkdir --dir $TESTDIR --method $METHOD $@ > $OUTPUT
		tail -n 10 $OUTPUT
	done
	echo
	echo
	echo
done

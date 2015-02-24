#!/usr/bin/env python

from sklearn.datasets import load_iris
from sklearn.ensemble import RandomForestClassifier

# the sample iris dataset contains 150 records on four dimensions
# classifiying into one of three results
# (the four dimensions happen to be sepal length, sepal width, petal length and
#  petal width, and the result happens to be which species of iris the given
#  flower belongs to, but the dataset is just numeric)
data = load_iris()

# For educational purposes, print one line per record, of the form
# [input1, ...inputN] result
for i, record in enumerate(data['data']):
    print record, data['target'][i]

# create and train a classifier on the sample data
classifier = RandomForestClassifier()
classifier.fit(data['data'], data['target'])

# Now we can use it to predict an input sample (which I shamelessly stole from
# the sample data, because I'm lazy)
print "Predicting an input sample:"
test = [6.7, 3.3, 5.7, 2.5]
print test, classifier.predict(test)

all: functional

env: requirements.txt
	rm -rf env
	virtualenv env
	env/bin/pip install -r requirements.txt

check: env
	env/bin/flake8 --exclude=env/ --max-line-length=120 -j 1
	#pylint --rcfile=pylintrc *.py

functional: env check
	. env/bin/activate && nosetests -v

clean:
	rm -rf env
	rm -f *.pyc store-service.log
.PHONY: functional check clean

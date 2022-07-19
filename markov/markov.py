import markovify
import sys
import flask

with open(sys.argv[1]) as f:
    text = f.read()
text_model = markovify.Text(text)

app = flask.Flask(__name__)


@app.route('/sentence')
def sentence():
    return text_model.make_sentence()


app.run()

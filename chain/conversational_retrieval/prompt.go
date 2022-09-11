package conversational_retrieval

// FIXME : we need a way to make KB lookup more generic
var instruction = `INSTRUCTION
You have access to this knowledge base (KB) :
name | description
wikipedia | knowledge base to find any general information, I accept standalone keywords as query.

CONVERSATION
{{.history}}
User: {{.question}}

When responding to me, please output a response in one of two formats:

**Option 1:**
use this if you can answer directly without KB lookup
Answer in this following json schema:

{
"conversation_context":"longer additional context to understand user 
# AI and NLP Design

## Objectives

- extract medications and dosage instructions
- extract investigations
- isolate clinical observations
- recover gracefully when AI output is imperfect

## Implementation Strategy

The backend uses an AI provider abstraction:

- `mock` provider
  - deterministic phrase and catalog matching
  - always available for local demos
- `openai` provider
  - OpenAI-compatible chat completion call
  - JSON-only output contract

## Reliability Measures

- normalized casing and whitespace
- inferred medication quantity when omitted
- reference pricing lookup after extraction
- structured parse metadata with provider, model, confidence, and warnings

## Why This Approach Works for the Assessment

- It satisfies the mandatory AI integration requirement.
- It remains runnable without external dependencies.
- It demonstrates production thinking by separating provider concerns from validation, billing, and persistence.

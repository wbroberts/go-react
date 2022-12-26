# Make React File

Wanted to learn Go so started with a simple CLI to generate react files to reduce the boilerplate.

## Install

`go install github.com/wbroberts/mrf`

## Use

`mrf`

## Component

Creates a component and component test file.

`mrf component <ComponentName>`

Options:

- `--props` or `-p` adds an empty props type for component
- `--skip-tests` skips adding the test file
- `--dir` or `-d` sets the path for the directory the component should be added to. Will create the directory if it does not exist yet

### Example:

`mrf component Example`

- Component `Example.component.tsx`

```tsx
import React from 'react';

export const Example = () => {
	return <div>Example renders</div>;
}
```

- Component test `Example.component.test.tsx`

```tsx
import { screen, render } from '@testing-library/react';

import { Example } from './Example.component';

describe('Example', () => {
	it('renders', () => {
		render(<Example />);

		expect(screen.getByText(/Example renders/)).toBeDefined();
	});
});
```

## Optional Config File for Defaults

Can add a config file to set default flags when running commands.

`touch make-react-file.yaml`

```yaml
# All values are optional
component:
  dir: src/components
  props: true
  skip-tests: false
```

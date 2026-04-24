---
trigger: always_on
---

# UI Frontend Rules

Use these rules as a baseline when working on the frontend, especially on a React, TypeScript, Vite, Inertia, Next.js Pages Router, or Tailwind stack.

## Basic Principles
- The UI should be clear, quick to understand, and feel intentionally designed.
- Avoid generic looks that feel like raw templates.
- Each page should have a clear visual hierarchy: heading, content area, action area, and feedback state.
- Ensure the layout remains comfortable on both mobile and desktop.
- If there isn't a more specific brand direction, use Sandi's default visual preferences: a blue and white base with a dark mode that remains elegant and legible.
- The default visual direction should tend toward clear, glossy blur, slight gradients, matte, brighter, and engaging without feeling cluttered.

## React, Next.js, And Component Preference
- Use TypeScript as the default, not plain JavaScript, for new React projects.

- If the frontend stack is React or Next.js, the default preferred UI base is `shadcn/ui`.
- `shadcn/ui` is used as a customizable foundation, not as a raw default.
- Reusable components should have clear and unambiguous API props.
- Don't make a single component handle table, filter, form, modal, and fetch logic all at once if it becomes difficult to maintain.
- For state, prioritize the simplest things first: local state, form state, and then shared state only if necessary.

## Blade Component Preference
- If the frontend uses Blade, prioritize custom Blade components built in-house.
- Use Tailwind for styling and Alpine.js for lightweight interactions when needed.
- Basic components like buttons, modals, tabs, inputs, selects, dropdowns, alerts, table wrappers, and cards should be built as reusable internal Blade components.
- Avoid heavy reliance on the Blade UI kit if your actual needs can be met with clean internal components.

## Visual Character Default
- The appearance should feel clean, modern, and light.
- Use bright, clean, and uncluttered surfaces for light mode.
- Use blur or glassy effects in a controlled manner for headers, panels, modals, dropdowns, or specific cards.
- Use subtle gradients for accents, hero panels, highlight panels, or identity elements, not to aggressively color the entire UI.
- Maintain a matte and professional feel; gloss should only be a visual accent, not an overly dramatic effect.
- Avoid UIs that are too dense, dark, dull, or heavy-looking.

## UI/UX Minimum
- All important pages should have a loading state.
- All forms should have clear validation feedback.
- All async requests should have a user-visible error state.
- Empty states should be informative, not empty pages devoid of context.
- The main action should be clearly visible without cluttering the UI.

## Default Color Direction
- Default base colors: blue and white.
- Light mode should feel clean, bright, professional, and credible.
- Dark mode should be customized, not simply turning the background black.
- Use blue for the primary action, highlight, focus ring, link, and main identity elements.
- Use white or near-white as the primary surface in light mode.
- For dark mode, use navy or dark slate blue as the primary surface with text that maintains comfortable contrast.

## Default Theme Tokens
If the user doesn't have a design system yet, use token directions like these as a baseline:

### Light Mode
- `background`: white or very pale blue
- `surface`: white
- `surface-muted`: very light gray-blue
- `primary`: bold, clean blue
- `primary-hover`: slightly darker blue
- `text`: dark slate or dark navy
- `text-muted`: medium gray-blue
- `border`: light gray-blue
- `gradient-accent`: gradient from light blue to white or light blue to soft cyan
- `success`, `warning`, `danger`: keep it semantic, don't force the blue

### Dark Mode
- `background`: dark navy or dark blue-slate
- `surface`: dark blue slightly lighter than the background
- `surface-muted`: muted slate blue
- `primary`: bright blue that remains strong on dark surfaces
- `primary-hover`: Lighter or more intense blue
- `text`: soft white or bright blue-white
- `text-muted`: light gray-blue
- `border`: dark blue-gray with sufficient contrast
- `gradient-accent`: gradient from navy to soft electric blue or dark blue to slate blue

## Styling Direction
- Use a consistent system of spacing, radius, shadow, and typography.
- If using Tailwind, keep the design system small: color, size, button variants, input, card, badge, alert.
- Use blur, transparency, and gradients subtly and measuredly.
- Avoid color combinations that are too low in contrast.
- Don't be biased towards a generic purple theme.
- If dark mode is used, maintain depth with graduated surfaces, not just one flat dark color.
- If there is no brand guide, choose a visual direction that is clean, modern, bright, matte, and credible.

## Form & Data Display
- Long forms should be broken down into sections.
- Tables should remain usable on small screens; Use stacking, card rows, or horizontal scrolling where comfortable.
- Number, date, currency, and status formats should be consistent throughout the application.
- Filters and search should be placed close to the data they affect.

## Basic Components That Should Be Consistent
- Button
- Input
- Select
- Textarea
- Modal/Dialog
- Tabs
- Table/DataTable
- Badge/Status
- Alert/Toast
- Empty state
- Loading skeleton or spinner
- Theme toggle if light/dark mode is available

## Accessibility Baseline
- Use clear form labels.
- Ensure buttons and links are easily distinguishable.
- Pay attention to the user's keyboard focus state.
- Don't rely solely on color to convey status/errors.
- Ensure the contrast between light and dark modes remains legible.

## Frontend Quality Gate
- No large, unmaintainable components.
- No async state that sits idle upon failure.
- No forms without validation and feedback.
- No critical UI elements that are great on desktop but break on mobile.
- No new UI dependencies without a clear technical or delivery rationale.
- No dark mode implementations that compromise readability or visual hierarchy.
- No blur, gloss, or gradient effects that degrade usability.